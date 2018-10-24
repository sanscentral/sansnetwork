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

	"github.com/sanscentral/sansnetwork/typeconv"
)

const (
	// Mainnet protocol message header bytes (0xF9BEB4D9)
	mainnetMagicBytes = "\xF9\xBE\xB4\xD9"

	// Testnet (3) protocol message header bytes (0x0B110907)
	testnetMagicBytes = "\x0B\x11\x09\x07"

	headerlen = 24
)

// Header is structure of BTC message header
type Header struct {
	Start      [4]byte  //char[4]
	Command    [12]byte //char[12]
	PayloadLen [4]byte  //uint32_t
	Checksum   [4]byte  //char[4]
	Full       []byte
}

var magicBytes []byte

// HeaderLength returns BTC protocol header length
func HeaderLength() int {
	return headerlen
}

// MagicBytes returns signal bytes
func MagicBytes(testnet bool) []byte {
	t := []byte(mainnetMagicBytes)
	if testnet {
		t = []byte(testnetMagicBytes)
	}
	return t
}

func (h *Header) verifyPayload(payload []byte) bool {
	chk := typeconv.CheckSumFromBytes(payload)
	return h.Checksum == chk
}

func makeHeader(command string, payload []byte, testnet bool) []byte {
	cmd := typeconv.CommandFromBytes(command)
	ln := typeconv.BytesFromUint32(uint32(len(payload)))
	chk := typeconv.CheckSumFromBytes(payload)

	header := []byte{}
	header = append(header, MagicBytes(testnet)...)
	header = append(header, cmd[:]...)
	header = append(header, ln[:]...)
	header = append(header, chk[:]...)
	return header
}

// ReadHeader reads and parses header directly from TCP connection
func ReadHeader(conn net.Conn) (Header, error) {
	responseHeader := make([]byte, headerlen)
	_, err := conn.Read(responseHeader)
	if err != nil {
		return Header{}, err
	}
	return ParseHeader(responseHeader)
}

// ParseHeader returns decoded message structure for given header
func ParseHeader(b []byte) (Header, error) {
	n := Header{}
	if len(b) != headerlen {
		return Header{}, errors.New("Invalid header length")
	}
	copy(n.Start[:], b[:4])
	copy(n.Command[:], b[4:16])
	copy(n.PayloadLen[:], b[16:20])
	copy(n.Checksum[:], b[20:24])
	n.Full = b
	return n, nil
}
