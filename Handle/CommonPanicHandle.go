package Handle

import (
	"MQTTStorage/Panic"
	"github.com/PharbersDeveloper/bp-go-lib/log"
	"net/http"
)

type CommonPanicHandle struct {
}

func (ctm CommonPanicHandle) NewCommonPanicHandle(args ...interface{}) CommonPanicHandle {
	return CommonPanicHandle{}
}

func (ctm CommonPanicHandle) HandlePanic(rw http.ResponseWriter, r *http.Request, p interface{}) {
	log.NewLogicLoggerBuilder().Build().Error("CommonHandlePanic接收到", p)
	Panic.ErrInstance().ErrorReval(p.(string), rw)
}