package msg

import (
	"io"

	"github.com/wuyazero/Elastos.ELA.Utility/p2p"
)

type VerAck struct{}

func (msg *VerAck) CMD() string {
	return p2p.CmdVerAck
}

func (msg *VerAck) Serialize(io.Writer) error {
	return nil
}

func (msg *VerAck) Deserialize(io.Reader) error {
	return nil
}
