package msg

import (
	"fmt"
	"io"

	"github.com/wuyazero/Elastos.ELA.Utility/common"
	"github.com/wuyazero/Elastos.ELA.Utility/p2p"
)

type GetData struct {
	InvList []*InvVect
}

func NewGetData() *GetData {
	msg := &GetData{
		InvList: make([]*InvVect, 0, defaultInvListSize),
	}
	return msg
}

// AddInvVect adds an inventory vector to the message.
func (msg *GetData) AddInvVect(iv *InvVect) error {
	if len(msg.InvList)+1 > MaxInvPerMsg {
		return fmt.Errorf("GetData.AddInvVect too many invvect in message [max %v]", MaxInvPerMsg)
	}

	msg.InvList = append(msg.InvList, iv)
	return nil
}

func (msg *GetData) CMD() string {
	return p2p.CmdGetData
}

func (msg *GetData) Serialize(writer io.Writer) error {
	count := uint32(len(msg.InvList))
	if err := common.WriteElement(writer, count); err != nil {
		return err
	}

	for _, vect := range msg.InvList {
		if err := vect.Serialize(writer); err != nil {
			return err
		}
	}

	return nil
}

func (msg *GetData) Deserialize(reader io.Reader) error {
	var count uint32
	if err := common.ReadElement(reader, &count); err != nil {
		return err
	}

	msg.InvList = make([]*InvVect, 0, count)
	for i := uint32(0); i < count; i++ {
		vect := new(InvVect)
		if err := vect.Deserialize(reader); err != nil {
			return err
		}
		msg.InvList = append(msg.InvList, vect)
	}

	return nil
}
