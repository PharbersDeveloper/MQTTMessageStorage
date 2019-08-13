package Strategy

import (
	"encoding/json"
	"fmt"
	"MQTTStorage/Daemons"
	"MQTTStorage/Model"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	emitter "github.com/emitter-io/go/v2"
	"github.com/go-redis/redis"
)

// Deprecated: 属于过度设计封禁
type MAXKafkaMessageStrategy struct {
	Rd *BmRedis.BmRedis
	Em *Daemons.Emitter
	//URI string
}

func (msms *MAXKafkaMessageStrategy) onMessageHandler(c *emitter.Client, msg emitter.Message) {
	// fmt.Printf("[emitter] -> [B] received on specific handler: '%s' topic: '%s'\n", msg.Payload(), msg.Topic())
	// 从Emitter的调试上来看，这个MessageHandler没用到，But这是必须的参数
}

func (msms *MAXKafkaMessageStrategy) DoExecute(msg Model.Message) (interface{}, error) {
	payload, _ := msg.PayLoad.(map[string]interface{})
	channel := payload["channel"].(string)

	//builder := &Builder.EmitterClientBuilder{}
	//director := &Builder.Director {Bud: builder}
	//emitterClient := director.Create(msms.URI, msms.onMessageHandler)
	//client := emitterClient.GetClient()

	rdClient := msms.Rd.GetRedisClient()
	result, err := rdClient.Get(fmt.Sprint("mqtt_channel_key_", channel)).Result()
	if err != redis.Nil {
		b, _ := json.Marshal(payload)
		client := msms.Em.GetClient()
		err = client.Publish(result, channel, b, emitter.WithAtLeastOnce())
	}
	return nil, err
}