package msg

import (
	"bytes"
	"encoding/binary"
	. "github.com/elastos/Elastos.ELA.Utility/common"
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

func (msg *Addrs) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := WriteElements(buf, uint32(len(msg.Addrs)), msg.Addrs)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (msg *Addrs) Deserialize(body []byte) error {
	buf := bytes.NewReader(body)
	count, err := ReadUint32(buf)
	if err != nil {
		return err
	}

	msg.Addrs = make([]Addr, count)
	err = binary.Read(buf, binary.LittleEndian, &msg.Addrs)
	if err != nil {
		return err
	}

	return nil
}
