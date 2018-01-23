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

package typeconv

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
)

// Uint32FromBytes converts bytes slice to uint32 type
func Uint32FromBytes(bytes []byte) uint32 {
	return binary.LittleEndian.Uint32(bytes)
}

// Uint64FromBytes converts byte slice to uint64 type
func Uint64FromBytes(bytes []byte) uint64 {
	return binary.LittleEndian.Uint64(bytes)
}

// Uint8FromBytes converts bytes to uint8 type
func Uint8FromBytes(bytes []byte) uint8 {
	if len(bytes) > 0 {
		return bytes[0]
	}
	return 0
}

// CheckSumFromBytes computes a twice iterated SHA256 4 bytes array from given slice
func CheckSumFromBytes(payload []byte) [4]byte {
	h := sha256.Sum256(payload)
	hb := sha256.Sum256(h[:])
	res := [4]byte{}
	for index := 0; index < 4; index++ {
		if index <= len(hb)-1 {
			res[index] = hb[index]
		}
	}
	return res
}

// CommandFromBytes returns command type array from string
func CommandFromBytes(command string) [12]byte {
	res := [12]byte{}
	b := []byte(command)
	for index := 0; index < 12; index++ {
		if index <= len(b)-1 {
			res[index] = b[index]
		}
	}
	return res
}

// Char16FromString returns char16 type from string
func Char16FromString(command string) [16]byte {
	res := [16]byte{}
	b := []byte(command)
	for index := 0; index < 16; index++ {
		if index <= len(b)-1 {
			res[index] = b[index]
		}
	}
	return res
}

// BytesFromInt32 returns byte array for int32 type
func BytesFromInt32(i int32) [4]byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, i)
	res := [4]byte{}
	for index := 0; index < 4; index++ {
		if index <= len(buf.Bytes())-1 {
			res[index] = buf.Bytes()[index]
		}
	}
	return res
}

// BytesFromUint32 returns byte array for uint32 type
func BytesFromUint32(i uint32) [4]byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, i)
	res := [4]byte{}
	for index := 0; index < 4; index++ {
		if index <= len(buf.Bytes())-1 {
			res[index] = buf.Bytes()[index]
		}
	}
	return res
}

// BytesFromUint16 returns byte array for uint16 type
func BytesFromUint16(i uint16) [2]byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, i)
	res := [2]byte{}
	for index := 0; index < 2; index++ {
		if index <= len(buf.Bytes())-1 {
			res[index] = buf.Bytes()[index]
		}
	}
	return res
}

// BytesFromUint64 returns byte array for uint64 type
func BytesFromUint64(i uint64) [8]byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, i)
	res := [8]byte{}
	for index := 0; index < 8; index++ {
		if index <= len(buf.Bytes())-1 {
			res[index] = buf.Bytes()[index]
		}
	}
	return res
}

// BytesFromInt64 returns byte array for int64 type
func BytesFromInt64(i int64) [8]byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, i)
	res := [8]byte{}
	for index := 0; index < 8; index++ {
		if index <= len(buf.Bytes())-1 {
			res[index] = buf.Bytes()[index]
		}
	}
	return res
}
