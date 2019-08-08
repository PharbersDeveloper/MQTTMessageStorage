package Strategy

import (
	"errors"
	"fmt"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/alfredyang1986/blackmirror/bmlog"
)

type header struct {
	Application	string	`json:"application"`
	Type		string	`json:"type"`
	MsgType		string	`json:"msgType"`
}

type Message struct {
	Header 	header	`json:"header"`
	Body 	interface{}		`json:"body"`
}

type MessageContext struct {
	strategy MessageStrategy
	Msg Message
	Rd *BmRedis.BmRedis
	URI string
}

func (mc *MessageContext) mapping() error {
	var err error
	switch mc.Msg.Header.MsgType {
	case "KeyGen":
		mc.strategy = &GenerateChannelKeyStrategy{Rd: mc.Rd, URI: mc.URI}
	case "RetrievingChannel":
		mc.strategy = &RetrievingChannelStrategy{Rd: mc.Rd, URI: mc.URI}
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
