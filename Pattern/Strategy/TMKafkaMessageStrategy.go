package Strategy

import (
	"github.com/PharbersDeveloper/MQTTMessageStorage/Daemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/alfredyang1986/blackmirror/bmlog"
	emitter "github.com/emitter-io/go/v2"
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

func (tkms *TMKafkaMessageStrategy) DoExecute(msg Message) (interface{}, error) {
	bmlog.StandardLogger().Info("TMKafkaMessageStrategy DoExecute")

	body := msg.Body.(map[string]interface{})
	channelKey := body["channelKey"].(string)
	channel := body["channel"].(string)


	//builder := &Builder.EmitterClientBuilder{}
	//director := &Builder.Director {Bud: builder}
	//emitterClient := director.Create(tkms.URI, tkms.onMessageHandler)
	//client := emitterClient.GetClient()

	client := tkms.Em.GetClient()

	err := client.Publish(channelKey, channel, body, emitter.WithAtLeastOnce())

	return nil, err
}
