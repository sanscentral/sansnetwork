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

	"github.com/sanscentral/sansnetwork/inventory"
	"github.com/sanscentral/sansnetwork/typeconv"
)

const (
	entryLen = 36
	typeLen  = 4
	countlen = 1
)

// ParseInventoryPayload parses byte payload into inventory structure
func ParseInventoryPayload(b []byte) (inventory.Item, error) {
	n := inventory.Item{}
	n.Entry = []inventory.Entry{}
	n.Count = typeconv.Uint8FromBytes(b[:countlen])

	for index := 0; index < int(n.Count); index++ {
		s := (entryLen * index) + countlen
		bytes := b[s:(s + typeLen)]
		new := inventory.Entry{}
		e := (s + entryLen)
		new.Type = typeconv.Uint32FromBytes(bytes)
		if len(b) < e {
			return n, errors.New("invalid inv payload length")
		}
		copy(new.Hash[:], b[(s+typeLen):e])
		n.Entry = append(n.Entry, new)
	}

	return n, nil
}
