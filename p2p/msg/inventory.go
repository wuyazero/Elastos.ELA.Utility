package msg

import (
	"io"

	. "github.com/elastos/Elastos.ELA.Utility/common"
)

type Inventory struct {
	Type   uint8
	Hashes []*Uint256
}

func NewInventory(dataType uint8, hashes []*Uint256) *Inventory {
	msg := new(Inventory)
	msg.Type = dataType
	msg.Hashes = hashes
	return msg
}

func (msg *Inventory) CMD() string {
	return "inv"
}

func (msg *Inventory) Serialize(writer io.Writer) error {
	count := uint32(len(msg.Hashes))
	return WriteElements(writer, msg.Type, count, msg.Hashes)
}

func (msg *Inventory) Deserialize(reader io.Reader) error {
	var count uint32
	err := ReadElements(reader, &msg.Type, &count)
	if err != nil {
		return err
	}

	msg.Hashes = make([]*Uint256, count)
	return ReadElement(reader, &msg.Hashes)
}
