package msg

import (
	"io"

	"github.com/wuyazero/Elastos.ELA.Utility/common"
	"github.com/wuyazero/Elastos.ELA.Utility/p2p"
)

type MerkleBlock struct {
	Header       common.Serializable
	Transactions uint32
	Hashes       []*common.Uint256
	Flags        []byte
}

func NewMerkleBlock(header common.Serializable) *MerkleBlock {
	return &MerkleBlock{Header: header}
}

func (msg *MerkleBlock) CMD() string {
	return p2p.CmdMerkleBlock
}

func (msg *MerkleBlock) Serialize(writer io.Writer) error {
	err := msg.Header.Serialize(writer)
	if err != nil {
		return err
	}

	return common.WriteElements(writer,
		msg.Transactions,
		uint32(len(msg.Hashes)),
		msg.Hashes, msg.Flags)
}

func (msg *MerkleBlock) Deserialize(reader io.Reader) error {
	err := msg.Header.Deserialize(reader)
	if err != nil {
		return err
	}

	msg.Transactions, err = common.ReadUint32(reader)
	if err != nil {
		return err
	}

	hashes, err := common.ReadUint32(reader)
	if err != nil {
		return err
	}

	msg.Hashes = make([]*common.Uint256, hashes)
	return common.ReadElements(reader, &msg.Hashes, &msg.Flags)
}
