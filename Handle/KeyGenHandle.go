package Handle

import (
	"MQTTStorage/Daemons"
	"MQTTStorage/Model"
	"MQTTStorage/Pattern/Strategy"
	"encoding/json"
	"github.com/PharbersDeveloper/bp-go-lib/log"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"reflect"
)

type KenGenHandler struct {
	Method     	string
	HttpMethod 	string
	Args       	[]string
	rd         	*BmRedis.BmRedis
	em			*Daemons.Emitter
}

func (k KenGenHandler) NewKenGenHandler(args ...interface{}) KenGenHandler {
	var r *BmRedis.BmRedis
	var em *Daemons.Emitter
	var hm string
	var md string
	var ag []string
	for i, arg := range args {
		if i == 0 {
			sts := arg.([]BmDaemons.BmDaemon)
			for _, dm := range sts {
				tp := reflect.ValueOf(dm).Interface()
				tm := reflect.ValueOf(tp).Elem().Type()
				if tm.Name() == "BmRedis" {
					r = dm.(*BmRedis.BmRedis)
				} else if tm.Name() == "Emitter" {
					em = dm.(*Daemons.Emitter)
				}
			}
		} else if i == 1 {
			md = arg.(string)
		} else if i == 2 {
			hm = arg.(string)
		} else if i == 3 {
			lst := arg.([]string)
			for _, str := range lst {
				ag = append(ag, str)
			}
		} else {
		}
	}

	return KenGenHandler{Method: md, HttpMethod: hm, Args: ag, rd: r, em: em}
}

func (k KenGenHandler) GetHttpMethod() string {
	return k.HttpMethod
}

func (k KenGenHandler) GetHandlerMethod() string {
	return k.Method
}

func (k KenGenHandler) KeyGen(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	w.Header().Add("Content-Type", "application/json")

	var response map[string]interface{}
	response = make(map[string]interface{})
	enc := json.NewEncoder(w)

	ERROR := func() int {
		response["status"] = "error"
		response["code"] = http.StatusInternalServerError
		response["msg"] = "Generate KeyGen Error"
		_ = enc.Encode(response)
		return 1
	}

	SUCCESS := func(res interface{}) int {
		response["status"] = "success"
		response["code"] = http.StatusOK
		response["msg"] = "Generate KeyGen Success"
		response["body"] = map[string]interface{}{ "channelKey": res.(string) }
		_ = enc.Encode(response)
		return 0
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil { log.NewLogicLoggerBuilder().Build().Error("MQTT读取参数出错 => ", err); return ERROR() }
	msg := Model.Message{}
	err = json.Unmarshal(body, &msg)
	if err != nil { log.NewLogicLoggerBuilder().Build().Error(err); return ERROR() }
	context := Strategy.MessageContext{ Msg: msg, Rd: k.rd, Em: k.em }
	res, err := context.DoExecute()
	if err != nil { log.NewLogicLoggerBuilder().Build().Error("生成MQTT Channel Key错误 => ", err); return ERROR() }
	return SUCCESS(res)
}
