package Builder

import (
	"MQTTStorage/Model"
	emitter "github.com/emitter-io/go/v2"
)

// Deprecated: 属于过度设计封禁
type Builder interface {
	SetURI(uri string) Builder
	SetMessageHandler(fun func(c *emitter.Client, msg emitter.Message)) Builder
	CreateClient() Builder
	Build() *Model.EmitterClient
}
