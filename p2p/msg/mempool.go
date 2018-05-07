package msg

import (
	"github.com/elastos/Elastos.ELA.Utility/p2p"
	"io"
)

type MemPool struct{}

func (msg *MemPool) CMD() string {
	return p2p.CmdMemPool
}

func (msg *MemPool) Serialize(io.Writer) error {
	return nil
}

func (msg *MemPool) Deserialize(io.Reader) error {
	return nil
}
