package bloom

import (
	"bytes"

	"github.com/elastos/Elastos.ELA.Utility/common"
)

type FilterLoad struct {
	Filter    []byte
	HashFuncs uint32
	Tweak     uint32
}

func (msg *FilterLoad) CMD() string {
	return "filterload"
}

func (msg *FilterLoad) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := common.WriteElements(buf, msg.Filter, msg.HashFuncs, msg.Tweak)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (msg *FilterLoad) Deserialize(body []byte) error {
	buf := bytes.NewReader(body)
	err := common.ReadElements(buf, &msg.Filter, &msg.HashFuncs, &msg.Tweak)
	if err != nil {
		return err
	}

	return nil
}
