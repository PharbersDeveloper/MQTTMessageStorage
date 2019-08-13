package Strategy

import (
	"MQTTStorage/Daemons"
	"MQTTStorage/Model"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/alfredyang1986/blackmirror/bmkafka"
	"github.com/alfredyang1986/blackmirror/bmlog"
	"github.com/elodina/go-avro"
	emitter "github.com/emitter-io/go/v2"
	"github.com/go-redis/redis"
	kafkaAvro "github.com/elodina/go-kafka-avro"
)

type RetrievingConsumerStrategy struct {
	Rd *BmRedis.BmRedis
	Em *Daemons.Emitter
}

func (rcs *RetrievingConsumerStrategy) onConsumerHandler(content interface{}) {
	c, _ := bmkafka.GetConfigInstance()

	decoder := kafkaAvro.NewKafkaAvroDecoder(c.SchemaRepositoryUrl)
	decoded, _ := decoder.Decode(content.([]byte))
	decodedRecord, _ := decoded.(*avro.GenericRecord)
	message := []byte(decodedRecord.String())

	msgModel := Model.Message{}
	json.Unmarshal(message, &msgModel)
	bmlog.StandardLogger().Warn(msgModel.Header.Channel)
	rdClient := rcs.Rd.GetRedisClient()
	result, err := rdClient.Get(fmt.Sprint("mqtt_channel_key_", msgModel.Header.Channel)).Result()
	if err != redis.Nil {
		client := rcs.Em.GetClient()
		err := client.Publish(result, msgModel.Header.Channel, fmt.Sprint(msgModel.PayLoad), emitter.WithAtLeastOnce())
		if err != nil { bmlog.StandardLogger().Error(err.Error()); panic(err.Error()) }
	}
}

func (rcs *RetrievingConsumerStrategy) DoExecute(msg Model.Message) (interface{}, error) {
	var err error
	payload := msg.PayLoad.(map[string]interface{})
	topic := payload["topic"].(string)
	if len(topic) > 0 {
		kafka, err := bmkafka.GetConfigInstance()
		kafka.Topics = []string{topic}
		if err == nil { go func() { kafka.SubscribeTopics(kafka.Topics, rcs.onConsumerHandler) }() }
	} else {
		err = errors.New("Topic Is Null")
	}

	return nil, err
}
