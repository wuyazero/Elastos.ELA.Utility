package p2p

import (
	"io"
)

// The message flying in the peer to peer network
type Message interface {
	// Get the message CMD parameter which is the type of this message
	CMD() string
	// Serialize the message content
	Serialize(io.Writer) error
	// Deserialize the message content through bytes
	Deserialize(io.Reader) error
}
