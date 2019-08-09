package Strategy

import (
	"github.com/PharbersDeveloper/MQTTMessageStorage/Daemons"
	"github.com/PharbersDeveloper/MQTTMessageStorage/Model"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/alfredyang1986/blackmirror/bmlog"
	emitter "github.com/emitter-io/go/v2"
)

type RetrievingConsumerStrategy struct {
	Rd *BmRedis.BmRedis
	Em *Daemons.Emitter
}

// TODO 感觉现在没啥用，先空着
func (rcs *RetrievingConsumerStrategy) onConsumerHandler(msg interface{}) {
	client := rcs.Em.GetClient()
	err := client.Publish("", "", msg, emitter.WithAtLeastOnce())
	if err != nil { bmlog.StandardLogger().Error(err.Error()); panic(err.Error()) }
}

func (rcs *RetrievingConsumerStrategy) DoExecute(msg Model.Message) (interface{}, error) {
	//var err error
	//body := msg.Body.(map[string]interface{})
	//channelKey := body["channelKey"].(string)
	//channel := body["channel"].(string)
	//
	//kafka, _ := bmkafka.GetConfigInstance()
	//kafka.SubscribeTopics()

	return nil, nil
}
