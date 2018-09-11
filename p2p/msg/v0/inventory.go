package v0

import (
	"io"

	"github.com/wuyazero/Elastos.ELA.Utility/common"
	"github.com/wuyazero/Elastos.ELA.Utility/p2p"
)

type Inv struct {
	Hashes []*common.Uint256
}

func NewInv(hashes []*common.Uint256) *Inv {
	return &Inv{Hashes: hashes}
}

func (msg *Inv) CMD() string {
	return p2p.CmdInv
}

func (msg *Inv) Serialize(w io.Writer) error {
	return common.WriteElements(w, uint32(len(msg.Hashes)), msg.Hashes)
}

func (msg *Inv) Deserialize(r io.Reader) error {
	count, err := common.ReadUint32(r)
	if err != nil {
		return err
	}

	msg.Hashes = make([]*common.Uint256, count)
	return common.ReadElement(r, &msg.Hashes)
}
