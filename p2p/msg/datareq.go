package msg

import (
	"io"

	. "github.com/elastos/Elastos.ELA.Utility/common"
)

type DataReq struct {
	Type uint8
	Hash Uint256
}

func NewDataReq(invType uint8, hash Uint256) *DataReq {
	dataReq := new(DataReq)
	dataReq.Type = invType
	dataReq.Hash = hash
	return dataReq
}

func (msg *DataReq) CMD() string {
	return "getdata"
}

func (msg *DataReq) Serialize(writer io.Writer) error {
	return WriteElements(writer, msg.Type, msg.Hash)
}

func (msg *DataReq) Deserialize(reader io.Reader) error {
	return ReadElements(reader, &msg.Type, &msg.Hash)
}
