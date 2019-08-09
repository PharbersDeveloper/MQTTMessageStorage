package Strategy

import (
	"errors"
	"fmt"
	"github.com/PharbersDeveloper/MQTTMessageStorage/Daemons"
	"github.com/PharbersDeveloper/MQTTMessageStorage/Model"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/alfredyang1986/blackmirror/bmlog"
)

type MessageContext struct {
	strategy MessageStrategy
	Msg Model.Message
	Rd *BmRedis.BmRedis
	Em *Daemons.Emitter
	//URI string
}

func (mc *MessageContext) mapping() error {
	var err error
	switch mc.Msg.Header.MsgType {
	case "KeyGen":
		mc.strategy = &GenerateChannelKeyStrategy{Rd: mc.Rd, Em: mc.Em}
	case "RetrievingChannel":
		mc.strategy = &RetrievingChannelStrategy{Rd: mc.Rd, Em: mc.Em}
	default:
		err = errors.New(fmt.Sprint(mc.Msg.Header.MsgType, "is not implementation"))
	}
	return err
}

func (mc *MessageContext) DoExecute() (interface{}, error) {
	err := mc.mapping()
	if err != nil {
		bmlog.StandardLogger().Error(err)
		return nil, err
	}
	return mc.strategy.DoExecute(mc.Msg)
}
