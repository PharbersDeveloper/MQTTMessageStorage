package Daemons

import (
	emitter "github.com/emitter-io/go/v2"
)

type Emitter struct {
	Host     string
	client *emitter.Client
}

func (e Emitter) NewEmitterDaemon(args map[string]string) *Emitter {
	em := &Emitter{Host: args["host"]}
	c, err := emitter.Connect(em.Host, e.onMessageHandle)

	if err != nil { panic(err.Error())}
	em.client = c

	return em
}


func (e *Emitter) onMessageHandle(c *emitter.Client, msg emitter.Message)  {
	// fmt.Printf("[emitter] -> [B] received on specific handler: '%s' topic: '%s'\n", msg.Payload(), msg.Topic())
	// 从Emitter的调试上来看，这个MessageHandler暂时没用到，But这是必须的参数
}

func (e *Emitter) GetClient() *emitter.Client {
	return e.client
}






