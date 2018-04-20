package msg

import (
	"bytes"
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

func (msg *Inventory) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	count := uint32(len(msg.Hashes))
	err := WriteElements(buf, msg.Type, count, msg.Hashes)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (msg *Inventory) Deserialize(body []byte) error {
	buf := bytes.NewReader(body)
	var count uint32
	err := ReadElements(buf, &msg.Type, &count)
	if err != nil {
		return err
	}

	msg.Hashes = make([]*Uint256, count)
	return ReadElement(buf, &msg.Hashes)
}
