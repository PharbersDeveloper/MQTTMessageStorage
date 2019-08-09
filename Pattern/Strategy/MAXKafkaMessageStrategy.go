package Strategy

import (
	"github.com/PharbersDeveloper/MQTTMessageStorage/Daemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/alfredyang1986/blackmirror/bmlog"
	emitter "github.com/emitter-io/go/v2"
)

type MAXKafkaMessageStrategy struct {
	Rd *BmRedis.BmRedis
	Em *Daemons.Emitter
	//URI string
}

func (msms *MAXKafkaMessageStrategy) onMessageHandler(c *emitter.Client, msg emitter.Message) {
	// fmt.Printf("[emitter] -> [B] received on specific handler: '%s' topic: '%s'\n", msg.Payload(), msg.Topic())
	// 从Emitter的调试上来看，这个MessageHandler没用到，But这是必须的参数
}

func (msms *MAXKafkaMessageStrategy) DoExecute(msg Message) (interface{}, error) {
	bmlog.StandardLogger().Info("MAXKafkaMessageStrategy DoExecute")

	body := msg.Body.(map[string]interface{})
	channelKey := body["channelKey"].(string)
	channel := body["channel"].(string)


	//builder := &Builder.EmitterClientBuilder{}
	//director := &Builder.Director {Bud: builder}
	//emitterClient := director.Create(msms.URI, msms.onMessageHandler)
	//client := emitterClient.GetClient()

	client := msms.Em.GetClient()

	err := client.Publish(channelKey, channel, body, emitter.WithAtLeastOnce())

	return nil, err
}