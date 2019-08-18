package test

import (
	"fmt"
	emitter "github.com/emitter-io/go/v2"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestMqttPublishTimeTask(t *testing.T) {
	t.Parallel()

	Convey("Test Mqtt Publish", t, func() {
		client, _ := emitter.Connect("tcp://123.56.179.133:46532", func(client *emitter.Client, message emitter.Message) {
			fmt.Printf("[emitter] -> [B] received: '%s' topic: '%s'\n", message.Payload(), message.Topic())
		})
		channelKey  := "XsKflXovpPuCKy4rGlioYVC7h6N1uutu"

		message := []byte(`{
					  "uuid": "TcLoIAH6Nt3jS4ai",
					  "progress": 0.13,
					  "status" : 1
					} `)

		ticker := time.NewTicker(10* time.Second)
		var ch chan int
		go func() {
			for range ticker.C {
				client.Publish(channelKey, "tm/", message, emitter.WithAtLeastOnce())
				fmt.Println("Oh My Good Lose Control")
			}
			ch <- 1
		}()
		<-ch
	})
}