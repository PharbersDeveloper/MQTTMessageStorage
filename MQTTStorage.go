package main

import (
	"MQTTStorage/Common/MQTTChannelState"
	"MQTTStorage/Factory"
	"net/http"
	"os"

	"github.com/PharbersDeveloper/bp-go-lib/log"
	"github.com/alfredyang1986/BmServiceDef/BmApiResolver"
	"github.com/alfredyang1986/BmServiceDef/BmConfig"
	"github.com/alfredyang1986/BmServiceDef/BmPodsDefine"
	"github.com/julienschmidt/httprouter"
	"github.com/manyminds/api2go"
)

func main() {

	//env.SetEnv()

	// 初始化MQTTChannel"伪池子"
	state := MQTTChannelState.StateSlice{}
	state.NewStateSlice()

	// 本地调试打开
	//os.Setenv("BM_KAFKA_CONF_HOME", fmt.Sprint(os.Getenv("BM_KAFKA_CONF_HOME"), "MQTTMessageStorage/resources/resource/kafkaconfig.json"))

	version := "v0"
	prodEnv := "MQTT_HOME"
	log.NewLogicLoggerBuilder().Build().Info("MQTT Server begins, version =", version)

	fac := Factory.Table{}
	var pod = BmPodsDefine.Pod{Name: "new MQTT", Factory: fac}
	envHome := os.Getenv(prodEnv)
	pod.RegisterSerFromYAML(envHome + "/resource/service-def.yaml")

	var bmRouter BmConfig.BmRouterConfig
	bmRouter.GenerateConfig(prodEnv)

	addr := bmRouter.Host + ":" + bmRouter.Port
	log.NewLogicLoggerBuilder().Build().Info("Listening on ", addr)
	api := api2go.NewAPIWithResolver(version, &BmApiResolver.RequestURL{Addr: addr})
	pod.RegisterAllResource(api)

	pod.RegisterAllFunctions(version, api)
	pod.RegisterAllMiddleware(api)

	handler := api.Handler().(*httprouter.Router)
	pod.RegisterPanicHandler(handler)
	http.ListenAndServe(":"+bmRouter.Port, handler)
}
