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

package message

import (
	"errors"
	"net"
	"time"

	"github.com/sanscentral/sansnetwork/seed"
	"github.com/sanscentral/sansnetwork/typeconv"
)

const (
	protocolVersion    = 70015              // Bitcoin Core 0.13.2
	services           = 0                  // No services supported on this node
	ip                 = "::ffff:127.0.0.1" // always return loopback
	nonce              = 0
	startHeight        = 1
	payloadLengthError = "Invalid payload length"
)

// Version structure for BTC version response
type Version struct {
	Version          [4]byte   // Highest protocol version understood by the transmitting node
	Services         [8]byte   // Services supported by the transmitting node
	Timestamp        [8]byte   // Current Unix epoch time
	RecieveServices  [8]byte   // Services supported by the receiving node
	RecieveIP        [16]byte  // IPv6 address of the receiving node
	RecievePort      [2]byte   // Port number of the receiving node
	TransmitServices [8]byte   // Services supported by the transmitting node
	TransmitIP       [16]byte  // IPv6 address of the transmitting node
	TransmitPort     [2]byte   // Port number of the transmitting node
	Nonce            [8]byte   // A random nonce which can help a node detect a connection to itself.
	UseragentLen     [1]byte   // Number of bytes in following user_agent field. (0)
	Useragent        [255]byte // User agent
	StartHeight      [4]byte   // The height of the transmitting node’s best block chain
	Relay            [1]byte   // bool. relay messages to this node?
	Full             []byte    // Raw full message bytes
}

// empty user agent results in smaller payload
var userAgent = []byte{00}

// ReadVersionPayload reads and parses version payload directly from TCP connection
func ReadVersionPayload(conn net.Conn, h Header) (Version, error) {
	len := typeconv.Uint32FromBytes(h.PayloadLen[:])
	vpl := make([]byte, len)
	_, err := conn.Read(vpl)
	if err != nil {
		return Version{}, err
	}
	return parseVersionPayload(vpl)
}

const (
	verPosEnd      = 4
	servPosEnd     = 12
	timeSmpPosEnd  = 20
	rcvSvcPosEnd   = 28
	rcvIPPosEnd    = 44
	rcvPrtPosEnd   = 46
	txSvcPosEnd    = 54
	txIPPosEnd     = 70
	txPrtPosEnd    = 72
	noncePosEnd    = 80
	uaLenPrtPosEnd = 81
	startHeightLen = 4
	relayLen       = 1
)

func parseVersionPayload(b []byte) (Version, error) {
	n := Version{}

	// Set size items
	if len(b) < uaLenPrtPosEnd {
		return Version{}, errors.New(payloadLengthError)
	}
	copy(n.Version[:], b[:verPosEnd])
	copy(n.Services[:], b[verPosEnd:servPosEnd])
	copy(n.Timestamp[:], b[servPosEnd:timeSmpPosEnd])
	copy(n.RecieveServices[:], b[timeSmpPosEnd:rcvSvcPosEnd])
	copy(n.RecieveIP[:], b[rcvSvcPosEnd:rcvIPPosEnd])
	copy(n.RecievePort[:], b[rcvIPPosEnd:rcvPrtPosEnd])
	copy(n.TransmitServices[:], b[rcvPrtPosEnd:txSvcPosEnd])
	copy(n.TransmitIP[:], b[txSvcPosEnd:txIPPosEnd])
	copy(n.TransmitPort[:], b[txIPPosEnd:txPrtPosEnd])
	copy(n.Nonce[:], b[txPrtPosEnd:noncePosEnd])
	copy(n.UseragentLen[:], b[noncePosEnd:uaLenPrtPosEnd])

	// Variable length
	varEnd := (uaLenPrtPosEnd + int(typeconv.Uint8FromBytes(n.UseragentLen[:])))
	if len(b) < varEnd {
		return Version{}, errors.New(payloadLengthError)
	}
	copy(n.Useragent[:], b[uaLenPrtPosEnd:varEnd])

	if len(b) < varEnd+startHeightLen {
		return Version{}, errors.New(payloadLengthError)
	}
	copy(n.StartHeight[:], b[varEnd:(varEnd+4)])
	varEnd += startHeightLen

	if len(b) < varEnd+relayLen {
		return Version{}, errors.New(payloadLengthError)
	}
	copy(n.Relay[:], b[varEnd:(varEnd+relayLen)])
	n.Full = b
	return n, nil
}

// NewVersionMessage creates a 'version' control message including header
func NewVersionMessage() []byte {
	payload := makeVersionPayload()
	header := makeHeader(CommandVersion, payload)
	return append(header, payload...)
}

// makeVersionMessage creates a 'version' payload
func makeVersionPayload() []byte {
	ver := typeconv.BytesFromInt32(protocolVersion)
	svc := typeconv.BytesFromUint64(services)
	tmstmp := typeconv.BytesFromInt64(time.Now().Unix())
	rsvc := typeconv.BytesFromUint64(services)
	addr := typeconv.Char16FromString(ip)
	prt := typeconv.BytesFromUint16(seed.MainnetPort)
	tsvc := typeconv.BytesFromUint64(services)
	taddr := typeconv.Char16FromString(ip)
	tprt := typeconv.BytesFromUint16(seed.MainnetPort)
	nonce := typeconv.BytesFromUint64(nonce)
	height := typeconv.BytesFromInt32(startHeight)
	rly := []byte{0x01} // Alway request relay

	payload := []byte{}
	payload = append(payload, ver[:]...)
	payload = append(payload, svc[:]...)
	payload = append(payload, tmstmp[:]...)
	payload = append(payload, rsvc[:]...)
	payload = append(payload, addr[:]...)
	payload = append(payload, prt[:]...)
	payload = append(payload, tsvc[:]...)
	payload = append(payload, taddr[:]...)
	payload = append(payload, tprt[:]...)
	payload = append(payload, nonce[:]...)
	payload = append(payload, userAgent...)
	payload = append(payload, height[:]...)
	payload = append(payload, rly[:]...)
	return payload
}
