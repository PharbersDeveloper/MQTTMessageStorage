package Model

import (
	emitter "github.com/emitter-io/go/v2"
)

// Deprecated: 属于过度设计封禁
type EmitterClient struct {
	URI string
	MessageHandler func(c *emitter.Client, msg emitter.Message)
	client *emitter.Client
}

func (ec *EmitterClient) SetURI(uri string) {
	ec.URI = uri
}

func (ec *EmitterClient) SetMessageHandler(fun func(c *emitter.Client, msg emitter.Message)) {
	ec.MessageHandler = fun
}

func (ec *EmitterClient) CreateClient() {
	c, _ := emitter.Connect(ec.URI, ec.MessageHandler)
	ec.client = c
}

func (ec *EmitterClient) GetURI() string {
	return ec.URI
}

func (ec *EmitterClient) GetClient() *emitter.Client {
	return ec.client
}

