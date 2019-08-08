package test

import (
	"fmt"
	emitter "github.com/emitter-io/go/v2"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGenerateKey(t *testing.T)  {
	Convey("Test Generate Key", t, func() {
		onMessageHandler := func(_ *emitter.Client, msg emitter.Message) {
			fmt.Printf("[emitter] -> [B] received on specific handler: '%s' topic: '%s'\n", msg.Payload(), msg.Topic())
		}

		// Create the client and connect to the broker
		client, _ := emitter.Connect("tcp://127.0.0.1:46532", onMessageHandler)

		// Set the presence handler
		client.OnPresence(func(_ *emitter.Client, ev emitter.PresenceEvent) {
			fmt.Printf("[emitter] -> [B] presence event: %d subscriber(s) at topic: '%s'\n", len(ev.Who), ev.Channel)
		})

		channelKey, _:= client.GenerateKey("Hv8HUCUDk6dFxoBttP7cp06UfEHzDXTU",
										"demo/test/",
									 "rwslp", 0)
		fmt.Println(channelKey)
		So(channelKey, ShouldNotEqual, "")
	})
}

