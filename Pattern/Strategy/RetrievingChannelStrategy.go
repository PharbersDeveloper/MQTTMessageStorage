package Strategy

import (
	"fmt"
	"github.com/PharbersDeveloper/MQTTMessageStorage/Common/MQTTChannelState"
	"github.com/PharbersDeveloper/MQTTMessageStorage/Daemons"
	"github.com/PharbersDeveloper/MQTTMessageStorage/Model"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	emitter "github.com/emitter-io/go/v2"
)

type RetrievingChannelStrategy struct {
	Rd *BmRedis.BmRedis
	Em *Daemons.Emitter
	//URI string
}

// TODO 向Kafka转发消息
func (rcs *RetrievingChannelStrategy) onMessageHandler(c *emitter.Client, msg emitter.Message) {
	fmt.Printf("RetrievingChannelStrategy => [emitter] -> [B] received on specific handler: '%s' topic: '%s'\n", msg.Payload(), msg.Topic())
}

func (rcs *RetrievingChannelStrategy) DoExecute(msg Model.Message) (interface{}, error) {
	var err error
	body := msg.PayLoad.(map[string]interface{})
	channelKey := body["channelKey"].(string)
	channel := body["channel"].(string)

	state := MQTTChannelState.StateSlice{}

	if !state.Exist(channel) {
		//builder := &Builder.EmitterClientBuilder{}
		//director := &Builder.Director {Bud: builder}
		//emitterClient := director.Create(rcs.URI, rcs.onMessageHandler)
		//client := emitterClient.GetClient()

		client := rcs.Em.GetClient()
		state.Push(channel)
		// 这边可能会有内存问题，压测试才知道
		go func() { err = client.SubscribeWithHistory(channelKey, channel, 1, rcs.onMessageHandler) }()
	}

	return nil, err
}
