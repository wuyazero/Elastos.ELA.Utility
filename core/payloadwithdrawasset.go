package core

import (
	"errors"
	"io"

	"github.com/elastos/Elastos.ELA.Utility/common"
)

const WithdrawAssetPayloadVersion byte = 0x00

type PayloadWithdrawAsset struct {
	BlockHeight uint32
}

func (t *PayloadWithdrawAsset) Data(version byte) []byte {
	return []byte{0}
}

func (t *PayloadWithdrawAsset) Serialize(w io.Writer, version byte) error {
	if err := common.WriteUint32(w, t.BlockHeight); err != nil {
		return errors.New("[WithdrawAsset], BlockHeight serialize failed.")
	}

	return nil
}

func (t *PayloadWithdrawAsset) Deserialize(r io.Reader, version byte) error {
	height, err := common.ReadUint32(r)
	if err != nil {
		return errors.New("[WithdrawAsset], BlockHeight deserialize failed.")
	}
	t.BlockHeight = height

	return nil
}
