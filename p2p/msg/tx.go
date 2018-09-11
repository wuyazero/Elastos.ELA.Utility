package msg

import (
	"io"

	"github.com/wuyazero/Elastos.ELA.Utility/common"
	"github.com/wuyazero/Elastos.ELA.Utility/p2p"
)

type Tx struct {
	Transaction common.Serializable
}

func NewTx(tx common.Serializable) *Tx {
	return &Tx{Transaction: tx}
}

func (msg *Tx) CMD() string {
	return p2p.CmdTx
}

func (msg *Tx) Serialize(writer io.Writer) error {
	return msg.Transaction.Serialize(writer)
}

func (msg *Tx) Deserialize(reader io.Reader) error {
	return msg.Transaction.Deserialize(reader)
}
