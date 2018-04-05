package transaction

import (
	"errors"

	"github.com/elastos/Elastos.ELA.Utility/core/transaction/payload"
)

var PayloadFactorySingleton PayloadFactory

//for different transaction types with different payload format
//and transaction process methods
type TransactionType byte

func (self TransactionType) Name() string {
	return PayloadFactorySingleton.Name(self)
}

const (
	CoinBase      TransactionType = 0x00
	RegisterAsset TransactionType = 0x01
	TransferAsset TransactionType = 0x02
	Record        TransactionType = 0x03
	Deploy        TransactionType = 0x04
)

type PayloadFactory interface {
	Create(txType TransactionType) (Payload, error)
	Name(txType TransactionType) string
}

type PayloadFactoryImpl struct {
}

func (factor *PayloadFactoryImpl) Name(txType TransactionType) string {
	switch txType {
	case CoinBase:
		return "CoinBase"
	case RegisterAsset:
		return "RegisterAsset"
	case TransferAsset:
		return "TransferAsset"
	case Record:
		return "Record"
	case Deploy:
		return "Deploy"
	default:
		return "Unknown"
	}
}

func (factor *PayloadFactoryImpl) Create(txType TransactionType) (Payload, error) {
	switch txType {
	case CoinBase:
		return new(payload.CoinBase), nil
	case RegisterAsset:
		return new(payload.RegisterAsset), nil
	case TransferAsset:
		return new(payload.TransferAsset), nil
	case Record:
		return new(payload.Record), nil
	case Deploy:
		return new(payload.DeployCode), nil
	default:
		return nil, errors.New("[Transaction], invalid transaction type.")
	}
}

func init() {
	PayloadFactorySingleton = &PayloadFactoryImpl{}
}
