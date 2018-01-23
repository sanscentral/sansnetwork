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

package seed

import (
	"math/rand"
	"net"
	"time"
)

const maxSeedAttempts = 10

// GetNodeIPs returns ip addresses from DNS seeds
func GetNodeIPs() ([]net.IP, error) {
	return getSeedNodes(0)
}

// TODO: testnet or mainnet?
func getSeedNodes(attempt int) ([]net.IP, error) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	host := mainnetDNSSeeds[r.Intn(len(mainnetDNSSeeds))]
	ip, err := net.LookupIP(host)
	if err != nil {
		if attempt >= maxSeedAttempts {
			return []net.IP{}, err
		}
		getSeedNodes(attempt + 1)
	}
	return ip, nil
}
