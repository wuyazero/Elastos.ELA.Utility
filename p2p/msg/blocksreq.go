package msg

import (
	"io"

	. "github.com/elastos/Elastos.ELA.Utility/common"
)

type BlocksReq struct {
	Locator  []*Uint256
	HashStop Uint256
}

func NewBlocksReq(locator []*Uint256, hashStop Uint256) *BlocksReq {
	blocksReq := new(BlocksReq)
	blocksReq.Locator = locator
	blocksReq.HashStop = hashStop
	return blocksReq
}

func (msg *BlocksReq) CMD() string {
	return "getblocks"
}

func (msg *BlocksReq) Serialize(writer io.Writer) error {
	return WriteElements(writer, uint32(len(msg.Locator)), msg.Locator, msg.HashStop)
}

func (msg *BlocksReq) Deserialize(reader io.Reader) error {
	count, err := ReadUint32(reader)
	if err != nil {
		return err
	}

	msg.Locator = make([]*Uint256, count)
	return ReadElements(reader, &msg.Locator, &msg.HashStop)
}
