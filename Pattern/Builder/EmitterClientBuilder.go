package Builder

import (
	"github.com/PharbersDeveloper/MQTTMessageStorage/Model"
	emitter "github.com/emitter-io/go/v2"
)

type EmitterClientBuilder struct {
	emitterClient *Model.EmitterClient
}

func (ecb *EmitterClientBuilder) SetURI(uri string) Builder {
	if ecb.emitterClient == nil {ecb.emitterClient = &Model.EmitterClient{}}
	ecb.emitterClient.SetURI(uri)
	return ecb
}

func (ecb *EmitterClientBuilder) SetMessageHandler(fun func(c *emitter.Client, msg emitter.Message)) Builder {
	if ecb.emitterClient == nil {ecb.emitterClient = &Model.EmitterClient{}}
	ecb.emitterClient.SetMessageHandler(fun)
	return ecb
}

func (ecb *EmitterClientBuilder) CreateClient() Builder {
	if ecb.emitterClient == nil {ecb.emitterClient = &Model.EmitterClient{}}
	if len(ecb.emitterClient.URI) == 0 || ecb.emitterClient.MessageHandler == nil {
		panic("URI Or MessageHandler is Nil")
	}
	ecb.emitterClient.CreateClient()
	return ecb
}

func (ecb *EmitterClientBuilder) Build() *Model.EmitterClient {
	return ecb.emitterClient
}