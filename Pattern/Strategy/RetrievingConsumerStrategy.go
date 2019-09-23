package Strategy

import (
	"MQTTStorage/Daemons"
	"MQTTStorage/Model"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PharbersDeveloper/bp-go-lib/log"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/alfredyang1986/blackmirror/bmkafka"
	"github.com/elodina/go-avro"
	kafkaAvro "github.com/elodina/go-kafka-avro"
	emitter "github.com/emitter-io/go/v2"
	"github.com/go-redis/redis"
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
	_ = json.Unmarshal(message, &msgModel)

	if len(msgModel.Header.Channel) > 0 {
		log.NewLogicLoggerBuilder().Build().Info("从Kafka中获取MQTT Channel地址  ", msgModel.Header.Channel)
		rdClient := rcs.Rd.GetRedisClient()
		result, err := rdClient.Get(fmt.Sprint("mqtt_channel_key_", msgModel.Header.Channel)).Result()
		if err != redis.Nil {
			client := rcs.Em.GetClient()
			err := client.Publish(result, msgModel.Header.Channel, fmt.Sprint(msgModel.PayLoad), emitter.WithAtLeastOnce())
			if err != nil {
				log.NewLogicLoggerBuilder().Build().Error("从Kafka发送获取MQTT Channel地址发送失败,错误信息  ",err)
				panic(err.Error())
			}
		} else if err != nil {
			log.NewLogicLoggerBuilder().Build().Error(err)
			panic(err)
		}
	}
}

func (rcs *RetrievingConsumerStrategy) DoExecute(msg Model.Message) (interface{}, error) {
	var err error
	payload := msg.PayLoad.(map[string]interface{})
	topic := payload["topic"].(string)
	if len(topic) > 0 {
		kafka, err := bmkafka.GetConfigInstance()
		kafka.Topics = []string{topic}
		if err == nil { go func() {
			log.NewLogicLoggerBuilder().Build().Info("Kafka 启动Consumer监听,Topic  ", topic)
			kafka.SubscribeTopics(kafka.Topics, rcs.onConsumerHandler)
		}() }
	} else {
		log.NewLogicLoggerBuilder().Build().Error("Kafka 启动Consumer监听失败，原因Topic为空!")
		err = errors.New("Topic Is Null")
	}

	return nil, err
}
