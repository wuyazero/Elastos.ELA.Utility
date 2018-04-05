package transaction

import (
	"bytes"
	"testing"

	"github.com/elastos/Elastos.ELA.Utility/core/transaction/payload"
)

//todo test others

func TestTransaction_DeserializeUnsignedWithoutType_PayloadTest(t *testing.T) {
	tx := Transaction{Payload: &payload.CoinBase{[]byte("ELA")}}
	buf := new(bytes.Buffer)
	err := tx.Serialize(buf)
	if err != nil {
		t.Error("Unexpect error")
	}
	err = tx.Deserialize(buf)
	if err != nil {
		t.Error("Unexpect error")
	}
	p, ok := tx.Payload.(*payload.CoinBase)
	if !ok {
		t.Error("Payload type error")
	}
	if string(p.CoinbaseData) != "ELA" {
		t.Error("Payload value error")
	}
}
