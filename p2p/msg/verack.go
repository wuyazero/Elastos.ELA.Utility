package msg

import (
	"github.com/elastos/Elastos.ELA.Utility/p2p"
	"io"
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
