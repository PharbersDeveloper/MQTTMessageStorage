package Builder

import (
	"github.com/PharbersDeveloper/MQTTMessageStorage/Model"
	emitter "github.com/emitter-io/go/v2"
)

type Director struct {
	Bud Builder
}

func (d Director) Create(uri string, handler func(c *emitter.Client, msg emitter.Message)) *Model.EmitterClient {
	return d.Bud.SetURI(uri).SetMessageHandler(handler).CreateClient().Build()
}
