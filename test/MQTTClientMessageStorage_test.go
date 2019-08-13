package test

import (
	"fmt"
	emitter "github.com/emitter-io/go/v2"
	. "github.com/smartystreets/goconvey/convey"
	"strconv"
	"testing"
	"time"
)

var onMessageHandler = func(_ *emitter.Client, msg emitter.Message) {
	fmt.Printf("[emitter] -> [B] received on specific handler: '%s' topic: '%s'\n", msg.Payload(), msg.Topic())
}

// Create the client and connect to the broker
var client, _ = emitter.Connect("tcp://127.0.0.1:46532", onMessageHandler)

func TestMQTTTopicToSendMessage(t *testing.T)  {
	Convey("Topic To Send Message", t, func() {
		var err error

		channelKey  := "UKKrMs2rhcHodW6KK57hOa47XB_VBowX"
		//channelKey  := "G239dRa72LGJMkqzsVRsI9ubr8T_xu1t"

		// Publish to the channel
		fmt.Println("[emitter] <- [B] publishing to 'demo/'")

		arrays := []int{1, 2, 3, 4, 5}
		for _, v := range arrays {
			err = client.Publish(channelKey, "demo/", strconv.Itoa(v), emitter.WithAtLeastOnce())
		}

		//arrays := []int{1, 2, 3, 4, 5}
		//for _, v := range arrays {
		//	err = client.Publish(channelKey, "test/", strconv.Itoa(v), emitter.WithAtLeastOnce())
		//}

		//err = client.Publish(channelKey, "demo/", "Fuck", emitter.WithAtLeastOnce())

		So(err, ShouldEqual, nil)
	})
}

func TestMQTTTopicToSendMessage2(t *testing.T)  {
	Convey("Topic To Send Message", t, func() {
		var err error

		channelKey  := "G239dRa72LGJMkqzsVRsI9ubr8T_xu1t"

		// Publish to the channel
		fmt.Println("[emitter] <- [B] publishing to 'demo/'")

		arrays := []int{6, 7, 8, 9, 10}
		for _, v := range arrays {
			err = client.Publish(channelKey, "test/", strconv.Itoa(v), emitter.WithAtLeastOnce())
		}

		//err = client.Publish(channelKey, "demo/", "Fuck", emitter.WithAtLeastOnce())

		So(err, ShouldEqual, nil)
	})
}

func TestMQTTTopicToReadMessage_1(t *testing.T)  {
	Convey("Read MQTT Message", t, func() {
		channelKey  := "UKKrMs2rhcHodW6KK57hOa47XB_VBowX"

		var callBackMessage string

		onMessageHandler := func(_ *emitter.Client, msg emitter.Message) {
			message := msg.Payload()
			fmt.Println("=================")
			fmt.Printf("[emitter] -> [B] received on specific handler: '%s' topic: '%s'\n", message, msg.Topic())
			callBackMessage = string(message)
		}
		fmt.Println(callBackMessage)
		// Create the client and connect to the broker
		//client, _ := emitter.Connect("tcp://127.0.0.1:46532", onMessageHandler)

		// Subscribe to demo channel
		fmt.Println("[emitter] <- [B] subscribing to 'demo/'")
		//_ = client.Subscribe(channelKey, "demo/", onMessageHandler)
		_ = client.SubscribeWithHistory(channelKey, "demo/",1, onMessageHandler)

		time.Sleep(180 * time.Second)

		So(callBackMessage, ShouldEqual, "Fuck")
	})
}

func TestMQTTTopicToReadMessage_2(t *testing.T)  {
	Convey("Read MQTT Message", t, func() {
		channelKey  := "8FhgCOzhbHH444urugqWBrRYY3bGI7J1"

		var callBackMessage string

		onMessageHandler := func(_ *emitter.Client, msg emitter.Message) {
			message := msg.Payload()
			fmt.Println("=================")
			fmt.Printf("[emitter] -> [B] received on specific handler: '%s' topic: '%s'\n", message, msg.Topic())
			callBackMessage = string(message)
		}
		fmt.Println(callBackMessage)

		// Create the client and connect to the broker
		//client, _ := emitter.Connect("tcp://127.0.0.1:46532", onMessageHandler)

		// Subscribe to demo channel
		fmt.Println("[emitter] <- [B] subscribing to 'test/'")
		//_ = client.Subscribe(channelKey, "demo/", onMessageHandler)
		_ = client.SubscribeWithHistory(channelKey, "test/",1, onMessageHandler)

		time.Sleep(180 * time.Second)

		//So(callBackMessage, ShouldEqual, "Fuck")
	})
}
