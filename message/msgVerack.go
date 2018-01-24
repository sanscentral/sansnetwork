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
	"net"

	"github.com/sanscentral/sansnetwork/typeconv"
)

// NewVerackMessage creates a message to acknowledge a previously-received version message
func NewVerackMessage(testnet bool) []byte {
	return makeHeader(CommandVersionAcknowledge, []byte(""), testnet)
}

// ReadVerackMessage reads verack message directly from TCP connection
func ReadVerackMessage(conn net.Conn) (bool, error) {
	h, err := ReadHeader(conn)
	if err != nil {
		return false, err
	}
	if typeconv.CleanStringFromBytes(h.Command[:]) == CommandVersionAcknowledge {
		return true, nil
	}
	return false, nil
}
