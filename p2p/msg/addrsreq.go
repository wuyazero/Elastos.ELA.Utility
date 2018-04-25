package msg

import "io"

type AddrsReq struct{}

func (msg *AddrsReq) CMD() string {
	return "getaddr"
}

func (msg *AddrsReq) Serialize(io.Writer) error {
	return nil
}

func (msg *AddrsReq) Deserialize(io.Reader) error {
	return nil
}
