package Factory

import (
	"MQTTStorage/Daemons"
	"MQTTStorage/Handle"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
)

type Table struct{}

var MAIL_MODEL_FACTORY = map[string]interface{}{
}

var MAIL_STORAGE_FACTORY = map[string]interface{}{
}

var MAIL_RESOURCE_FACTORY = map[string]interface{}{
}

var MAIL_FUNCTION_FACTORY = map[string]interface{}{
	"CommonPanicHandle":        Handle.CommonPanicHandle{},
	"KenGenHandler":            Handle.KenGenHandler{},
	"RetrievingChannelHandler": Handle.RetrievingChannelHandler{},
	"RetrievingConsumerHandler": Handle.RetrievingConsumerHandler{},
	"PublishHandler":			Handle.PublishHandler{},
}
var MAIL_MIDDLEWARE_FACTORY = map[string]interface{}{
}

var MAIL_DAEMON_FACTORY = map[string]interface{}{
	"BmMongodbDaemon": BmMongodb.BmMongodb{},
	"BmRedisDaemon":   BmRedis.BmRedis{},
	"EmitterDaemon":   Daemons.Emitter{},
}

func (t Table) GetModelByName(name string) interface{} {
	return MAIL_MODEL_FACTORY[name]
}

func (t Table) GetResourceByName(name string) interface{} {
	return MAIL_RESOURCE_FACTORY[name]
}

func (t Table) GetStorageByName(name string) interface{} {
	return MAIL_STORAGE_FACTORY[name]
}

func (t Table) GetDaemonByName(name string) interface{} {
	return MAIL_DAEMON_FACTORY[name]
}

func (t Table) GetFunctionByName(name string) interface{} {
	return MAIL_FUNCTION_FACTORY[name]
}

func (t Table) GetMiddlewareByName(name string) interface{} {
	return MAIL_MIDDLEWARE_FACTORY[name]
}
