package test

import (
	"fmt"
	emitter "github.com/emitter-io/go/v2"
	//emitter "github.com/eclipse/paho.mqtt.golang"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMQTTNewClient(t *testing.T)  {
	t.Parallel()
	Convey("Test New MQTT Client", t,  func() {
		client, _ := emitter.Connect("tcp://127.0.0.1:46532", func(client *emitter.Client, message emitter.Message) {
			fmt.Printf("[emitter] -> [B] received: '%s' topic: '%s'\n", message.Payload(), message.Topic())
		})


		Convey("Client is Connect", func() {

			So(client.IsConnected(), ShouldEqual, true)
		})


	})
}

