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

package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/sanscentral/sansnetwork"
	"github.com/sanscentral/sansnetwork/inventory"
)

func main() {
	// Create a new connection to the BTC network
	networkconn, err := sansnetwork.NewNetworkConnection()
	if err != nil {
		panic(err)
	}
	defer networkconn.Close()

	// Set callback to handle new inventory entry messages
	inventory.SetInventoryHandler(invHandler)

	// Leave connection running
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for range c {
			fmt.Printf("\nStopping...\n")
			os.Exit(0)
		}
	}()
	wg.Wait()
}

func invHandler(i []inventory.Entry) {
	fmt.Printf("Recieved %d inventory - First is: %s \n", len(i), i[0].HexString())
}
