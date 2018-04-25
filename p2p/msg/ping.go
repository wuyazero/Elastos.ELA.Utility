package msg

import (
	"encoding/binary"
	"io"
)

type Ping struct {
	Height uint64
}

func NewPing(height uint32) *Ping {
	ping := new(Ping)
	ping.Height = uint64(height)
	return ping
}

func (msg *Ping) CMD() string {
	return "ping"
}

func (msg *Ping) Serialize(writer io.Writer) error {
	return binary.Write(writer, binary.LittleEndian, msg.Height)
}

func (msg *Ping) Deserialize(reader io.Reader) error {
	return binary.Read(reader, binary.LittleEndian, &msg.Height)
}
