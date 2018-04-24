package msg

import "io"

type VerAck struct{}

func (msg *VerAck) CMD() string {
	return "verack"
}

func (msg *VerAck) Serialize(io.Writer) error {
	return nil
}

func (msg *VerAck) Deserialize(io.Reader) error {
	return nil
}
