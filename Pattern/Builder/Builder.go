package Builder

import (
	"github.com/PharbersDeveloper/MQTTMessageStorage/Model"
	emitter "github.com/emitter-io/go/v2"
)

type Builder interface {
	SetURI(uri string) Builder
	SetMessageHandler(fun func(c *emitter.Client, msg emitter.Message)) Builder
	CreateClient() Builder
	Build() *Model.EmitterClient
}
