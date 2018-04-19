package msg

import (
	"bytes"

	"github.com/elastos/Elastos.ELA.Utility/core"
)

type Block struct {
	core.Block
}

func NewBlock(block core.Block) *Block {
	return &Block{Block: block}
}

func (msg *Block) CMD() string {
	return "block"
}

func (msg *Block) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := msg.Block.Serialize(buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (msg *Block) Deserialize(body []byte) error {
	return msg.Block.Deserialize(bytes.NewReader(body))
}
