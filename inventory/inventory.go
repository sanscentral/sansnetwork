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

package inventory

import "encoding/hex"

// Handler is type for inventory handler func
type Handler func([]Entry)

var handlerInstance Handler

// Item is inventory structure of BTC payload
type Item struct {
	Count uint8
	Entry []Entry
}

// Entry is a single inventory entry
type Entry struct {
	Type uint32
	Hash [32]byte
}

// HexString returns entry hash as a hex encoded string
func (e *Entry) HexString() string {
	return hex.EncodeToString(e.Hash[:])
}

// CallHandler calls set handler
func CallHandler(m []Entry) {
	if handlerInstance != nil {
		handlerInstance(m)
	}
}

// SetInventoryHandler for inventory entry callbacks
func SetInventoryHandler(i Handler) {
	handlerInstance = i
}
