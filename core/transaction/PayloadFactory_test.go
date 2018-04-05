package transaction

import (
	"testing"

	p "github.com/elastos/Elastos.ELA.Utility/core/transaction/payload"
)

func TestPayloadFactoryImpl_Name(t *testing.T) {
	name := PayloadFactorySingleton.Name(CoinBase)
	if name != "CoinBase" {
		t.Errorf("TransactionTypeName: [%v], actually: [%v]", "CoinBase", name)
	}

	name = PayloadFactorySingleton.Name(RegisterAsset)
	if name != "RegisterAsset" {
		t.Errorf("TransactionTypeName: [%v], actually: [%v]", "RegisterAsset", name)
	}

	name = PayloadFactorySingleton.Name(TransferAsset)
	if name != "TransferAsset" {
		t.Errorf("TransactionTypeName: [%v], actually: [%v]", "TransferAsset", name)
	}

	name = PayloadFactorySingleton.Name(Record)
	if name != "Record" {
		t.Errorf("TransactionTypeName: [%v], actually: [%v]", "Record", name)
	}

	name = PayloadFactorySingleton.Name(Deploy)
	if name != "Deploy" {
		t.Errorf("TransactionTypeName: [%v], actually: [%v]", "Deploy", name)
	}

	name = PayloadFactorySingleton.Name(0x05)
	if name != "Unknown" {
		t.Errorf("TransactionTypeName: [%v], actually: [%v]", "Unknown", name)
	}
}

func TestPayloadFactoryImpl_Create(t *testing.T) {
	payload, err := PayloadFactorySingleton.Create(CoinBase)
	if _, ok := payload.(*p.CoinBase); !ok {
		t.Error("Payload create error.")
	}
	if err != nil {
		t.Error("Unexpect error.")
	}

	payload, err = PayloadFactorySingleton.Create(RegisterAsset)
	if _, ok := payload.(*p.RegisterAsset); !ok {
		t.Error("Payload create error.")
	}
	if err != nil {
		t.Error("Unexpect error.")
	}

	payload, err = PayloadFactorySingleton.Create(TransferAsset)
	if _, ok := payload.(*p.TransferAsset); !ok {
		t.Error("Payload create error.")
	}
	if err != nil {
		t.Error("Unexpect error.")
	}

	payload, err = PayloadFactorySingleton.Create(Record)
	if _, ok := payload.(*p.Record); !ok {
		t.Error("Payload create error.")
	}
	if err != nil {
		t.Error("Unexpect error.")
	}

	payload, err = PayloadFactorySingleton.Create(Deploy)
	if _, ok := payload.(*p.DeployCode); !ok {
		t.Error("Payload create error.")
	}
	if err != nil {
		t.Error("Unexpect error.")
	}

	payload, err = PayloadFactorySingleton.Create(0x05)
	if payload != nil {
		t.Error("Payload create error.")
	}
	if err == nil {
		t.Error("Expect an error.")
	}
}
