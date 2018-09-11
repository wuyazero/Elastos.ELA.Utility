package v0

import (
	"io"

	"github.com/wuyazero/Elastos.ELA.Utility/common"
	"github.com/wuyazero/Elastos.ELA.Utility/p2p"
)

type NotFound struct {
	Hash common.Uint256
}

func NewNotFound(hash common.Uint256) *NotFound {
	return &NotFound{Hash: hash}
}

func (msg *NotFound) CMD() string {
	return p2p.CmdNotFound
}

func (msg *NotFound) Serialize(w io.Writer) error {
	return common.WriteElement(w, msg.Hash)
}

func (msg *NotFound) Deserialize(r io.Reader) error {
	return common.ReadElement(r, &msg.Hash)
}
