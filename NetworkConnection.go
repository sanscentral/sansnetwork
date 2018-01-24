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

// NetworkConnection is a self-managing connection to the bitcoin network
type NetworkConnection struct {
	nodes              []node.Connection
	die                bool
	testnet            bool
	requestedNodeCount int
}

// NewNetworkConnection starts a new connection to the bitcoin network
func NewNetworkConnection(nodeCount int, testnet bool) (NetworkConnection, error) {
	newc := NetworkConnection{
		testnet:            testnet,
		requestedNodeCount: nodeCount,
	}
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
// this Network Connection is connected to
func (c *NetworkConnection) NodeCount() int {
	return len(c.nodes)
}

// TODO: Periodically check if connections are alive and top-up as needed 'node.connected will be false for dead connections'
func (c *NetworkConnection) seedConnectionPool() {
	for len(c.nodes) < c.requestedNodeCount {
		if c.die {
			break
		}
		node, err := node.NewConnection(c.testnet)
		if err != nil {
			continue
		}
		c.nodes = append(c.nodes, node)
	}
}
