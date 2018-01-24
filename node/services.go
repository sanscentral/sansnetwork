package node

// ServiceFlag identifies services supported by a network peer.
type ServiceFlag uint64

const (
	// ServiceFullNode is a flag used to indicate a peer is a full node.
	ServiceFullNode ServiceFlag = 1 << iota

	// ServiceGetUTXO is a flag used to indicate a peer supports the getutxos and utxos commands (BIP0064).
	ServiceGetUTXO

	// ServiceBloom is a flag used to indicate a peer supports bloom filtering.
	ServiceBloom

	// ServiceWitness is a flag used to indicate a supports for blocks and transactions including witness data (BIP0144).
	ServiceWitness

	_

	// ServiceBCH indicates node is for bitcoin cash
	ServiceBCH
)
