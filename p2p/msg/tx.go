package msg

import (
	"bytes"

	"github.com/elastos/Elastos.ELA.Utility/core"
)

type Tx struct {
	core.Transaction
}

func NewTx(tx core.Transaction) *Tx {
	return &Tx{Transaction: tx}
}

func (msg *Tx) CMD() string {
	return "tx"
}

func (msg *Tx) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := msg.Transaction.Serialize(buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (msg *Tx) Deserialize(body []byte) error {
	return msg.Transaction.Deserialize(bytes.NewReader(body))
}
