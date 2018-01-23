/*
	SansNetwork is a  library for direct Bitcoin protocol interaction
	Copyright (C) 2018 Sans Central
	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Affero General Public License as
	published by the Free Software Foundation, either version 3 of the
	License, or (at your option) any later version.
	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU Affero General Public License for more details.
	You should have received a copy of the GNU Affero General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package node

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/sanscentral/sansnetwork/inventory"
	"github.com/sanscentral/sansnetwork/message"
	"github.com/sanscentral/sansnetwork/seed"
	"github.com/sanscentral/sansnetwork/typeconv"
)

const (
	initialConnectionTimeoutSec = 1
	maxConnectionAttempts       = 10
	pingDelaySec                = 5
	nonceVal                    = 78
)

var knownNodes []net.IP

// Connection is a single network node connection
type Connection struct {
	useragent    string
	host         string
	nonce        string
	services     uint64
	conn         net.Conn
	sendHeaders  bool
	handling     bool
	die          bool
	connected    bool
	ping         int64
	pendingPings map[uint64]int64
}

// Close single node connection
func (n *Connection) Close() {
	n.die = true
	n.conn.Close()
}

// UserAgent returns node useragent
func (n *Connection) UserAgent() string {
	return n.useragent
}

// startHandling indicates that a node event is currently being handled
func (n *Connection) startHandling() {
	n.handling = true
}

// endHandling indicates that a node event has finished being handled
func (n *Connection) endHandling() {
	n.handling = false
}

// startHeartBeat starts keep alive pings every X seconds
func (n *Connection) startHeartBeat() {
	n.performPing()
	for range time.Tick(pingDelaySec * time.Second) {
		if n.die {
			break
		}
		n.performPing()
	}
}

// performPing sends ping command to node
func (n *Connection) performPing() {
	if n.connected && !n.handling {
		nonce := time.Now().UnixNano() + nonceVal
		ping := message.NewPingMessage(uint64(nonce))
		_, err := n.conn.Write(ping)
		if err != nil {
			n.Close()
			n.connected = false
		}
		n.pendingPings[uint64(nonce)] = time.Now().UnixNano()
	}
}

// handle is the primay function for handling incoming node events
func (n *Connection) handle(b []byte) {
	n.startHandling()
	defer n.endHandling()
	if bytes.HasPrefix(b, message.MagicBytes()) {
		h, err := message.ParseHeader(b)
		var cmd string
		var payload []byte
		if err == nil {
			cmd = typeconv.CleanStringFromBytes(h.Command[:])
			payloadlength := typeconv.Uint32FromBytes(h.PayloadLen[:])
			if payloadlength > 0 {
				payload = make([]byte, payloadlength)
				n.conn.Read(payload)
			}
		} else {
			cmd = message.CommandError
		}

		fmt.Printf("Command: %s\n", h.Command)
		//fmt.Printf("Payload: %v\n", payload)

		// Handle command and payloads for this node
		switch cmd {
		case message.CommandSendHeaders:
			n.sendHeaders = true
		case message.CommandInventory:
			inv, _ := message.ParseInventoryPayload(payload)
			if len(inv.Entry) > 0 {
				inventory.CallHandler(inv.Entry)
			}
		case message.CommandPing:
			// Respond to ping with pong
			nonce := message.ReadPingPayload(payload)
			ping := message.NewPongMessage(nonce)
			n.conn.Write(ping)
		case message.CommandPong:
			// Recieved pong
			nonce := message.ReadPongPayload(payload)
			tm := ((time.Now().UnixNano() - n.pendingPings[nonce]) / int64(time.Millisecond))
			delete(n.pendingPings, nonce)
			n.ping = tm
		case message.CommandError:
			// Handle node error
		}
	}
}

// listen starts listening to the node
func (n *Connection) listen() {
	for {
		if n.die {
			break
		}
		if !n.handling {
			response := make([]byte, message.HeaderLength())
			n.conn.Read(response)
			if len(response) != 0 {
				n.handle(response)
			}
		}
	}
}

// NewConnection creates a single new node connection
func NewConnection() (Connection, error) {
	if len(knownNodes) == 0 {
		seeds, err := seed.GetNodeIPs()
		if err != nil {
			return Connection{}, errors.New("Failed to find a node")
		}
		knownNodes = append(knownNodes, seeds...)
	}

	s := rand.NewSource(time.Now().UnixNano())
	attempts := 0
	new := Connection{}
	var conn net.Conn
	var attemptedNode net.IP
	for {
		r := rand.New(s)
		attemptedNode = knownNodes[r.Intn(len(knownNodes))]
		serv := fmt.Sprintf("%s:%d", attemptedNode.String(), seed.MainnetPort)
		var err error
		conn, err = net.DialTimeout("tcp", serv, initialConnectionTimeoutSec*time.Second)
		if err == nil {
			break
		}
		attempts++
		if attempts >= maxConnectionAttempts {
			return Connection{}, errors.New("Cannot connect to node exceeded max attempts")
		}
	}

	var err error
	// Send version
	versionMsg := message.NewVersionMessage()
	_, err = conn.Write(versionMsg)
	if err != nil {
		conn.Close()
		removeFromKnownNodes(attemptedNode)
		return Connection{}, fmt.Errorf("Failed to write version: %s", err.Error())
	}

	// Recieve version
	header, err := message.ReadHeader(conn)
	if err != nil {
		conn.Close()
		removeFromKnownNodes(attemptedNode)
		return Connection{}, fmt.Errorf("Header not understood: %s", err.Error())
	}

	versionResponse, err := message.ReadVersionPayload(conn, header)
	if err != nil {
		conn.Close()
		removeFromKnownNodes(attemptedNode)
		return Connection{}, fmt.Errorf("Version message not understood: %s", err.Error())
	}

	// Send verack
	verackMsg := message.NewVerackMessage()
	_, err = conn.Write(verackMsg)
	if err != nil {
		conn.Close()
		removeFromKnownNodes(attemptedNode)
		return Connection{}, fmt.Errorf("Failed to write ver ack: %s", err.Error())
	}

	// Recieve verack
	ready, err := message.ReadVerackMessage(conn)
	if err != nil {
		conn.Close()
		removeFromKnownNodes(attemptedNode)
		return Connection{}, fmt.Errorf("Verack not understood: %s", err.Error())
	}
	if ready {
		// Connection success
		new.connected = true
	} else {
		conn.Close()
		removeFromKnownNodes(attemptedNode)
		return Connection{}, errors.New("Did not recieve verack where expected")
	}

	new.conn = conn
	new.useragent = typeconv.CleanStringFromBytes(versionResponse.Useragent[:])
	new.host = conn.RemoteAddr().String()
	new.nonce = fmt.Sprintf("%d", versionResponse.Nonce)
	new.services = typeconv.Uint64FromBytes(versionResponse.Services[:])
	new.pendingPings = map[uint64]int64{}

	fmt.Printf("Connected to:%s\n", new.UserAgent())

	go new.listen()
	go new.startHeartBeat()
	return new, nil
}

func removeFromKnownNodes(ip net.IP) {
	// TODO: remove attemptedNode from known nodes
}
