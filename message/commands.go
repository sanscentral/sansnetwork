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

const (
	// CommandVersion is the command used when a node creates an outgoing connection
	// it will immediately advertise its version
	// The remote node will respond with its version
	CommandVersion = "version"

	// CommandVersionAcknowledge message is sent in reply to version
	CommandVersionAcknowledge = "verack"

	// CommandSendHeaders is sent to request Direct headers announcement (instead of inv command)
	CommandSendHeaders = "sendheaders"

	// CommandInventory Allows a node to advertise its knowledge of one or more objects
	// It can be received unsolicited, or in reply to getblocks
	CommandInventory = "inv"

	// CommandPing is sent primarily to confirm that the TCP/IP connection is still valid
	CommandPing = "ping"

	// CommandPong  is sent in response to a CommandPing
	CommandPong = "pong"

	// CommandError is used to represent en erroneous command
	CommandError = "error"

	//******************TODO**********************************
	// Commands below this line are pending Implementation ***
	//******************TODO**********************************

	// CommandBlock is sent in response to a getdata message which requests transaction information from a block hash
	CommandBlock = "block"

	// CommandGetBlocks returns an inv packet containing the list of blocks starting right after the
	// last known hash in the block locator object, up to stop value or 500 blocks (max)
	CommandGetBlocks = "getblocks"

	// CommandGetData is used in response to inv, to retrieve the content of a specific object (usually sent after receiving an inv packet)
	CommandGetData = "getdata"

	// CommandGetHeaders return a headers packet containing the headers of blocks starting right after the
	// last known hash in the block locator object, up to stop value or 2000 blocks (max)
	CommandGetHeaders = "getheaders"

	// CommandHeaders returns block headers in response to a getheaders packet
	CommandHeaders = "headers"

	// CommandMempool asks for information about transactions a node has verified but which have not yet confirmed
	CommandMempool = "mempool"

	// CommandNotFound  is a response to CommandGetData if requested transaction was not in the memory pool or relay set
	CommandNotFound = "notfound"

	// CommandTx describes a bitcoin transaction, in response to CommandGetData
	CommandTx = "tx"

	// CommandAddress provides information on known nodes of the network.
	CommandAddress = "addr"

	// CommandFeeFilter is used to filter transaction invs for transactions that fall below the feerate provided in the CommandFeeFilter message interpreted as satoshis per KB
	CommandFeeFilter = "feefilter"

	// CommandGetAddress sends a request to a node asking for information about known active peers
	CommandGetAddress = "getaddr"

	// CommandReject is sent when messages are rejected.
	CommandReject = "reject"

	// CommandAlert provides alert messages (removed from bitcoin core in March 2016) : Deprecated
	CommandAlert = "alert"

	/****
	* Bloom filter commands
	*****/

	// CommandMerkleBlock is related to Bloom filtering of connections and is defined in BIP 0037
	CommandMerkleBlock = "merkleblock"
	// CommandFilterAdd is related to Bloom filtering of connections and is defined in BIP 0037
	CommandFilterAdd = "filteradd"
	// CommandFilterClear is related to Bloom filtering of connections and is defined in BIP 0037
	CommandFilterClear = "filterclear"
	// CommandFilterLoad is related to Bloom filtering of connections and is defined in BIP 0037
	CommandFilterLoad = "filterload"

	// ***
)
