package Strategy

import (
	"fmt"
	"github.com/PharbersDeveloper/MQTTMessageStorage/Pattern/Builder"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	emitter "github.com/emitter-io/go/v2"
	"github.com/go-redis/redis"
	"time"
)

type GenerateChannelKeyStrategy struct {
	Rd *BmRedis.BmRedis
	URI string
}

func (g *GenerateChannelKeyStrategy) onMessageHandler(c *emitter.Client, msg emitter.Message) {
	// fmt.Printf("[emitter] -> [B] received on specific handler: '%s' topic: '%s'\n", msg.Payload(), msg.Topic())
	// 从Emitter的调试上来看，这个MessageHandler没用到，But这是必须的参数
}

func (g *GenerateChannelKeyStrategy) DoExecute(msg Message) (interface{}, error) {
	body := msg.Body.(map[string]interface{})
	key := body["securityKey"].(string)
	channel := body["channel"].(string)
	permissions := body["permissions"].(string)
	ttl := int(body["ttl"].(float64))

	rdClient := g.Rd.GetRedisClient()
	defer rdClient.Close()
	result, err := rdClient.Get(fmt.Sprint("mqtt_channel_key_", channel)).Result()
	if err == redis.Nil {
		builder := &Builder.EmitterClientBuilder{}
		director := &Builder.Director {Bud: builder}
		emitterClient := director.Create(g.URI, g.onMessageHandler)
		client := emitterClient.GetClient()
		key, err := client.GenerateKey(key, channel, permissions, ttl)
		g.pushRedisData(fmt.Sprint("mqtt_channel_key_", channel), key, time.Duration(ttl) * time.Second)
		return key, err
	}
	return result, nil
}

func (g *GenerateChannelKeyStrategy) pushRedisData(key string, value interface{}, time time.Duration) error {
	rdClient := g.Rd.GetRedisClient()
	defer rdClient.Close()
	pipe := rdClient.Pipeline()
	pipe.Set(key, value, time)
	_, err := pipe.Exec()
	return err
}
