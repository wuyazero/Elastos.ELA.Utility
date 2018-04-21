package core

import (
	"errors"
	"io"

	. "github.com/elastos/Elastos.ELA.Utility/common"
)

const WithdrawAssetPayloadVersion byte = 0x00

type PayloadWithdrawAsset struct {
	BlockHeight              uint32
	GenesisBlockAddress      string
	SideChainTransactionHash string
}

func (t *PayloadWithdrawAsset) Data(version byte) []byte {
	return []byte{0}
}

func (t *PayloadWithdrawAsset) Serialize(w io.Writer, version byte) error {
	if err := WriteElements(w, t.BlockHeight, t.GenesisBlockAddress, t.SideChainTransactionHash); err != nil {
		return errors.New("[WithdrawAsset], BlockHeight serialize failed.")
	}
  
	return nil
}

func (t *PayloadWithdrawAsset) Deserialize(r io.Reader, version byte) error {
	return ReadElements(r, &t.BlockHeight, &t.GenesisBlockAddress, &t.SideChainTransactionHash)
}
