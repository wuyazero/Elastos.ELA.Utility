package msg

import (
	. "github.com/elastos/Elastos.ELA.Utility/common"
	"io"
)

type NotFound struct {
	Hash Uint256
}

func NewNotFound(hash Uint256) *NotFound {
	return &NotFound{Hash: hash}
}

func (msg *NotFound) CMD() string {
	return "notfound"
}

func (msg *NotFound) Serialize(writer io.Writer) error {
	return msg.Hash.Serialize(writer)
}

func (msg *NotFound) Deserialize(reader io.Reader) error {
	return msg.Hash.Deserialize(reader)
}
