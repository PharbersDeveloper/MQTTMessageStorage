package Builder

import (
	"MQTTStorage/Model"
	emitter "github.com/emitter-io/go/v2"
)

// Deprecated: 属于过度设计封禁
type Director struct {
	Bud Builder
}

func (d Director) Create(uri string, handler func(c *emitter.Client, msg emitter.Message)) *Model.EmitterClient {
	return d.Bud.SetURI(uri).SetMessageHandler(handler).CreateClient().Build()
}
