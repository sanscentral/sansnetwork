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

package sansnetwork

import (
	"github.com/sanscentral/sansnetwork/node"
)

// seedConnectionCount is the number of connected nodes from seed
const seedConnectionCount = 1

// NetworkConnection is a self-managing connection to the bitcoin network
type NetworkConnection struct {
	nodes []node.Connection
	die   bool
}

// NewNetworkConnection starts a new connection to the bitcoin network
// TODO: add options parameter to toggle testnet/mainnet, number of nodes,
// log level & Maintain existing connections (Check for dead)
func NewNetworkConnection() (NetworkConnection, error) {
	newc := NetworkConnection{}
	go newc.seedConnectionPool()
	return newc, nil
}

// Close connection to the bitcoin network
func (c *NetworkConnection) Close() {
	for _, n := range c.nodes {
		n.Close()
	}
}

// NodeCount returns the number of active nodes
// this Network Connection is DIRECTLY connected to
func (c *NetworkConnection) NodeCount() int {
	return len(c.nodes)
}

func (c *NetworkConnection) seedConnectionPool() {
	for index := 0; index < seedConnectionCount-1; index++ {
		if c.die {
			break
		}
		node, err := node.NewConnection()
		if err != nil {
			continue
		}
		c.nodes = append(c.nodes, node)
	}
}
