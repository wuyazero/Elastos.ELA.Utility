package core

import (
	"errors"
	"io"

	"github.com/elastos/Elastos.ELA.Utility/common"
)

type PayloadTransferCrossChainAsset struct {
	// string: publickey; uint64: output index
	AddressesMap map[string]uint64
}

func (a *PayloadTransferCrossChainAsset) Data(version byte) []byte {
	//TODO: implement TransferCrossChainAsset.Data()
	return []byte{0}
}

func (a *PayloadTransferCrossChainAsset) Serialize(w io.Writer, version byte) error {
	if a.AddressesMap == nil {
		return errors.New("Invalid publickey map")
	}

	if err := common.WriteVarUint(w, uint64(len(a.AddressesMap))); err != nil {
		return errors.New("publicKey map's length serialize failed")
	}

	for k, v := range a.AddressesMap {
		if err := common.WriteVarString(w, k); err != nil {
			return errors.New("publicKey map's key serialize failed")
		}

		if err := common.WriteVarUint(w, v); err != nil {
			return errors.New("publicKey map's value serialize failed")
		}
	}

	return nil
}

func (a *PayloadTransferCrossChainAsset) Deserialize(r io.Reader, version byte) error {
	if a.AddressesMap == nil {
		return errors.New("Invalid public key map")
	}

	length, err := common.ReadVarUint(r, 0)
	if err != nil {
		return errors.New("publicKey map's length deserialize failed")
	}

	a.AddressesMap = nil
	a.AddressesMap = make(map[string]uint64)
	for i := uint64(0); i < length; i++ {
		k, err := common.ReadVarString(r)
		if err != nil {
			return errors.New("publicKey map's key deserialize failed")
		}

		v, err := common.ReadVarUint(r, 0)
		if err != nil {
			return errors.New("publicKey map's value deserialize failed")
		}

		a.AddressesMap[k] = v
	}

	return nil
}
