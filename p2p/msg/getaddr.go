package msg

import (
	"github.com/elastos/Elastos.ELA.Utility/p2p"
	"io"
)

type GetAddr struct{}

func (msg *GetAddr) CMD() string {
	return p2p.CmdGetAddr
}

func (msg *GetAddr) Serialize(io.Writer) error {
	return nil
}

func (msg *GetAddr) Deserialize(io.Reader) error {
	return nil
}
