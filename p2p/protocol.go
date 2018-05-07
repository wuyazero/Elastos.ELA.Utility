package p2p

const (
	// EIP001Version is the protocol version support SPV messages.
	EIP001Version uint32 = 10001

	TxData    = 0x01
	BlockData = 0x02

	MaxHeaderHashes = 100
	MaxBlocksPerMsg = 500
)
