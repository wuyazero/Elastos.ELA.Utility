package p2p

import (
	"io"
	"bytes"
)

// The message flying in the peer to peer network
type Message interface {
	// Get the message CMD parameter which is the type of this message
	CMD() string
	// Serialize the message content
	Serialize(io.Writer)  error
	// Deserialize the message content through bytes
	Deserialize(io.Reader) error
}

// BuildMessage create the message header and return serialized message bytes
func BuildMessage(msg Message) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := msg.Serialize(buf)
	if err != nil {
		return nil, err
	}
	hdr, err := buildHeader(msg.CMD(), buf.Bytes()).Serialize()
	if err != nil {
		return nil, err
	}

	return append(hdr, buf.Bytes()...), nil
}
