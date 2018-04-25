package msg

import (
	"encoding/binary"
	. "github.com/elastos/Elastos.ELA.Utility/common"
	"io"
)

type Addrs struct {
	Addrs []Addr
}

func NewAddrs(addrs []Addr) *Addrs {
	msg := new(Addrs)
	msg.Addrs = addrs
	return msg
}

func (msg *Addrs) CMD() string {
	return "addr"
}

func (msg *Addrs) Serialize(writer io.Writer) error {
	return WriteElements(writer, uint32(len(msg.Addrs)), msg.Addrs)
}

func (msg *Addrs) Deserialize(reader io.Reader) error {
	count, err := ReadUint32(reader)
	if err != nil {
		return err
	}

	msg.Addrs = make([]Addr, count)
	return binary.Read(reader, binary.LittleEndian, &msg.Addrs)
}
