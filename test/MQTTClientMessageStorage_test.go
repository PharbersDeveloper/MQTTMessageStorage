package test

import (
	"fmt"
	emitter "github.com/emitter-io/go/v2"
	"os"
	"time"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMQTTTopicToSendMessage(t *testing.T)  {
	Convey("Topic To Send Message", t, func() {
		var err error
		channelKey  := os.Getenv("EMITTER_CHANNEL_KEY")

		onMessageHandler := func(_ *emitter.Client, msg emitter.Message) {
			fmt.Printf("[emitter] -> [B] received on specific handler: '%s' topic: '%s'\n", msg.Payload(), msg.Topic())
		}

		// Create the client and connect to the broker
		client, _ := emitter.Connect("tcp://127.0.0.1:46532", onMessageHandler)

		// Set the presence handler
		client.OnPresence(func(_ *emitter.Client, ev emitter.PresenceEvent) {
			fmt.Printf("[emitter] -> [B] presence event: %d subscriber(s) at topic: '%s'\n", len(ev.Who), ev.Channel)
		})

		// Publish to the channel
		fmt.Println("[emitter] <- [B] publishing to 'demo/'")
		err = client.Publish(channelKey, "demo/", "Fuck", emitter.WithAtLeastOnce())

		So(err, ShouldEqual, nil)
	})
}

func TestMQTTTopicToReadMessage_1(t *testing.T)  {
	Convey("Read MQTT Message", t, func() {
		channelKey  := os.Getenv("EMITTER_CHANNEL_KEY")

		var callBackMessage string

		onMessageHandler := func(_ *emitter.Client, msg emitter.Message) {
			message := msg.Payload()
			fmt.Println("=================")
			fmt.Printf("[emitter] -> [B] received on specific handler: '%s' topic: '%s'\n", message, msg.Topic())
			callBackMessage = string(message)
		}
		fmt.Println(callBackMessage)
		// Create the client and connect to the broker
		client, _ := emitter.Connect("tcp://127.0.0.1:46532", onMessageHandler)

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
		channelKey  := os.Getenv("EMITTER_CHANNEL_KEY")

		var callBackMessage string

		onMessageHandler := func(_ *emitter.Client, msg emitter.Message) {
			message := msg.Payload()
			fmt.Println("=================")
			fmt.Printf("[emitter] -> [B] received on specific handler: '%s' topic: '%s'\n", message, msg.Topic())
			callBackMessage = string(message)
		}
		fmt.Println(callBackMessage)

		// Create the client and connect to the broker
		client, _ := emitter.Connect("tcp://127.0.0.1:46532", onMessageHandler)

		// Subscribe to demo channel
		fmt.Println("[emitter] <- [B] subscribing to 'demo/'")
		//_ = client.Subscribe(channelKey, "demo/", onMessageHandler)
		_ = client.SubscribeWithHistory(channelKey, "demo/",1, onMessageHandler)

		time.Sleep(180 * time.Second)

		//So(callBackMessage, ShouldEqual, "Fuck")
	})
}
