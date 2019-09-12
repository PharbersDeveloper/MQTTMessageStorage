package Strategy

import (
	"MQTTStorage/Common/MQTTChannelState"
	"MQTTStorage/Daemons"
	"MQTTStorage/Model"
	"encoding/json"
	"fmt"
	"github.com/PharbersDeveloper/bp-go-lib/log"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/alfredyang1986/blackmirror/bmkafka"
	"github.com/elodina/go-avro"
	kafkaAvro "github.com/elodina/go-kafka-avro"
	emitter "github.com/emitter-io/go/v2"
	"github.com/go-redis/redis"
)

type RetrievingChannelStrategy struct {
	Rd *BmRedis.BmRedis
	Em *Daemons.Emitter
	//URI string
}

// TODO 向Kafka转发消息
func (rcs *RetrievingChannelStrategy) onMessageHandler(c *emitter.Client, msg emitter.Message) {
	message := Model.Message{}
	_ = json.Unmarshal(msg.Payload() ,&message)
	topic := message.Header.Topic
	log.NewLogicLoggerBuilder().Build().Infof("接收到MQTT Channel消息: '%s' MQTTTopic: '%s'\n", msg.Payload(), msg.Topic())

	if len(topic) > 0 {
		kafka, _ := bmkafka.GetConfigInstance()
		encoder := kafkaAvro.NewKafkaAvroEncoder(kafka.SchemaRepositoryUrl)
		schema, err := avro.ParseSchema(string(msg.Payload()))
		log.NewLogicLoggerBuilder().Build().Error(err)
		record := avro.NewGenericRecord(schema)
		log.NewLogicLoggerBuilder().Build().Error(err)
		recordByteArr, err := encoder.Encode(record)
		log.NewLogicLoggerBuilder().Build().Error(err)
		kafka.Produce(&topic, recordByteArr)
	}

}

func (rcs *RetrievingChannelStrategy) DoExecute(msg Model.Message) (interface{}, error) {
	var err error
	payload := msg.PayLoad.(map[string]interface{})
	//channelKey := payload["channelKey"].(string)
	channel := payload["channel"].(string)

	state := MQTTChannelState.StateSlice{}

	if !state.Exist(channel) {
		//builder := &Builder.EmitterClientBuilder{}
		//director := &Builder.Director {Bud: builder}
		//emitterClient := director.Create(rcs.URI, rcs.onMessageHandler)
		//client := emitterClient.GetClient()

		rdClient := rcs.Rd.GetRedisClient()
		result, err := rdClient.Get(fmt.Sprint("mqtt_channel_key_", channel)).Result()
		if err == redis.Nil || err != nil { log.NewLogicLoggerBuilder().Build().Error(err) }
		if err != redis.Nil {
			client := rcs.Em.GetClient()
			state.Push(channel)
			// 这边可能会有内存问题，压测试才知道
			go func() {
				err = client.SubscribeWithHistory(result, channel, 1, rcs.onMessageHandler)
				if err != nil { log.NewLogicLoggerBuilder().Build().Error(err) }
			}()

		}
	}

	return nil, err
}
