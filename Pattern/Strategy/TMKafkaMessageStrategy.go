package Strategy

import (
	"encoding/json"
	"fmt"
	"github.com/PharbersDeveloper/MQTTMessageStorage/Daemons"
	"github.com/PharbersDeveloper/MQTTMessageStorage/Model"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/alfredyang1986/blackmirror/bmlog"
	emitter "github.com/emitter-io/go/v2"
	"github.com/go-redis/redis"
)

type TMKafkaMessageStrategy struct {
	Rd *BmRedis.BmRedis
	Em *Daemons.Emitter
	//URI string
}

func (tkms *TMKafkaMessageStrategy) onMessageHandler(c *emitter.Client, msg emitter.Message) {
	// fmt.Printf("[emitter] -> [B] received on specific handler: '%s' topic: '%s'\n", msg.Payload(), msg.Topic())
	// 从Emitter的调试上来看，这个MessageHandler没用到，But这是必须的参数
}

func (tkms *TMKafkaMessageStrategy) DoExecute(msg Model.Message) (interface{}, error) {
	bmlog.StandardLogger().Info("TMKafkaMessageStrategy DoExecute")

	payload, _ := msg.PayLoad.(map[string]interface{})
	channel := payload["channel"].(string)


	//builder := &Builder.EmitterClientBuilder{}
	//director := &Builder.Director {Bud: builder}
	//emitterClient := director.Create(tkms.URI, tkms.onMessageHandler)
	//client := emitterClient.GetClient()

	rdClient := tkms.Rd.GetRedisClient()
	result, err := rdClient.Get(fmt.Sprint("mqtt_channel_key_", channel)).Result()
	if err != redis.Nil {
		b, _ := json.Marshal(payload)
		client := tkms.Em.GetClient()
		err = client.Publish(result, channel, b, emitter.WithAtLeastOnce())
	}

	return nil, err
}
