package Strategy

import (
	"fmt"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	//"github.com/alfredyang1986/blackmirror/bmkafka"
	emitter "github.com/emitter-io/go/v2"
)

type RetrievingConsumerStrategy struct {
	Rd *BmRedis.BmRedis
	URI string
}

// TODO 向Kafka转发消息
func (rcs *RetrievingConsumerStrategy) onMessageHandler(c *emitter.Client, msg emitter.Message) {
	fmt.Printf("RetrievingConsumerStrategy => [emitter] -> [B] received on specific handler: '%s' topic: '%s'\n", msg.Payload(), msg.Topic())
}

func (rcs *RetrievingConsumerStrategy) DoExecute(msg Message) (interface{}, error) {
	//var err error
	//body := msg.Body.(map[string]interface{})
	//channelKey := body["channelKey"].(string)
	//channel := body["channel"].(string)
	//
	//kafka, _ := bmkafka.GetConfigInstance()
	//kafka.SubscribeTopics()

	return nil, nil
}
