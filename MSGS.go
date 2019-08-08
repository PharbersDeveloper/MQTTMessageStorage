package main

import (
	"github.com/PharbersDeveloper/MQTTMessageStorage/Common/MQTTChannelState"
	"github.com/PharbersDeveloper/MQTTMessageStorage/Factory"
	"github.com/alfredyang1986/BmServiceDef/BmApiResolver"
	"github.com/alfredyang1986/BmServiceDef/BmConfig"
	"github.com/alfredyang1986/BmServiceDef/BmPodsDefine"
	"github.com/alfredyang1986/blackmirror/bmlog"
	"github.com/julienschmidt/httprouter"
	"github.com/manyminds/api2go"
	"net/http"
	"os"
)

func main() {
	// 初始化MQTTChannel"伪池子"
	state := MQTTChannelState.StateSlice{}
	state.NewStateSlice()

	version := "v0"
	prodEnv := "MQTT_HOME"
	bmlog.StandardLogger().Info("MQTT Server begins, version =", version)

	fac := Factory.Table{}
	var pod = BmPodsDefine.Pod{Name: "new MQTT", Factory: fac}
	envHome := os.Getenv(prodEnv)
	pod.RegisterSerFromYAML(envHome + "/resource/service-def.yaml")

	var bmRouter BmConfig.BmRouterConfig
	bmRouter.GenerateConfig(prodEnv)

	addr := bmRouter.Host + ":" + bmRouter.Port
	bmlog.StandardLogger().Info("Listening on ", addr)
	api := api2go.NewAPIWithResolver(version, &BmApiResolver.RequestURL{Addr: addr})
	pod.RegisterAllResource(api)

	pod.RegisterAllFunctions(version, api)
	pod.RegisterAllMiddleware(api)

	handler := api.Handler().(*httprouter.Router)
	pod.RegisterPanicHandler(handler)
	http.ListenAndServe(":"+bmRouter.Port, handler)

	bmlog.StandardLogger().Info("MQTT Server ends, version =", version)
}