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
	"github.com/sanscentral/sansnetwork/typeconv"
)

// NewPingMessage bytes for ping message including header
func NewPingMessage(nonce uint64, testnet bool) []byte {
	payload := typeconv.BytesFromUint64(nonce)
	header := makeHeader(CommandPing, payload[:], testnet)
	return append(header, payload[:]...)
}

// ReadPingPayload returns nonce for given ping
func ReadPingPayload(b []byte) uint64 {
	return typeconv.Uint64FromBytes(b)
}
