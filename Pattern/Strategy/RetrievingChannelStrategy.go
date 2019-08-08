package Strategy

import (
	"fmt"
	"github.com/PharbersDeveloper/MQTTMessageStorage/Pattern/Builder"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	emitter "github.com/emitter-io/go/v2"
)

type RetrievingChannelStrategy struct {
	Rd *BmRedis.BmRedis
	URI string
}

// TODO 向Kafka转发消息
func (rcs *RetrievingChannelStrategy) onMessageHandler(c *emitter.Client, msg emitter.Message) {
	fmt.Printf("RetrievingChannelStrategy => [emitter] -> [B] received on specific handler: '%s' topic: '%s'\n", msg.Payload(), msg.Topic())
}

func (rcs *RetrievingChannelStrategy) DoExecute(msg Message) (interface{}, error) {
	var err error
	body := msg.Body.(map[string]interface{})
	channelKey := body["channelKey"].(string)
	channel := body["channel"].(string)

	builder := &Builder.EmitterClientBuilder{}
	director := &Builder.Director {Bud: builder}
	emitterClient := director.Create(rcs.URI, rcs.onMessageHandler)
	client := emitterClient.GetClient()

	// 这边可能会有内存问题，压测试才知道
	go func() { err = client.SubscribeWithHistory(channelKey, channel, 1, rcs.onMessageHandler) }()

	return nil, err
}
