package msg

import (
	"bytes"

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

func (msg *BlocksReq) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := WriteElements(buf, uint32(len(msg.Locator)), msg.Locator, msg.HashStop)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (msg *BlocksReq) Deserialize(body []byte) error {
	buf := bytes.NewReader(body)
	count, err := ReadUint32(buf)
	if err != nil {
		return err
	}

	msg.Locator = make([]*Uint256, count)
	return ReadElements(buf, &msg.Locator, &msg.HashStop)
}
