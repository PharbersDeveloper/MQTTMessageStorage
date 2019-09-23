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

type RetrievingChannelHandler struct {
	Method     	string
	HttpMethod 	string
	Args       	[]string
	rd         	*BmRedis.BmRedis
	em			*Daemons.Emitter
}

func (r RetrievingChannelHandler) NewRetrievingChannelHandler(args ...interface{}) RetrievingChannelHandler {
	var rd *BmRedis.BmRedis
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
					rd = dm.(*BmRedis.BmRedis)
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

	return RetrievingChannelHandler{ Method: md, HttpMethod: hm, Args: ag, rd: rd, em: em }
}

func (r RetrievingChannelHandler) GetHttpMethod() string {
	return r.HttpMethod
}

func (r RetrievingChannelHandler) GetHandlerMethod() string {
	return r.Method

}

func (r RetrievingChannelHandler) RetrievingChannel(w http.ResponseWriter, req *http.Request, _ httprouter.Params) int {
	w.Header().Add("Content-Type", "application/json")
	var response map[string]interface{}
	response = make(map[string]interface{})
	enc := json.NewEncoder(w)

	ERROR := func() int {
		response["status"] = "error"
		response["code"] = http.StatusInternalServerError
		response["msg"] = "Channel监听失败"
		_ = enc.Encode(response)
		return 1
	}


	body, err := ioutil.ReadAll(req.Body)
	if err != nil { log.NewLogicLoggerBuilder().Build().Error("监听Channel读取参数出错  ", err); return ERROR() }
	msg := Model.Message{}
	err = json.Unmarshal(body, &msg)
	if err != nil { log.NewLogicLoggerBuilder().Build().Error("解析地址发送参数出错  ", err); return ERROR() }

	context := Strategy.MessageContext{ Msg: msg, Rd: r.rd, Em: r.em }
	_, err = context.DoExecute()
	if err != nil { log.NewLogicLoggerBuilder().Build().Error("监听Channel出错  ", err); return ERROR() }

	log.NewLogicLoggerBuilder().Build().Infof("MQTT 开启监听Channel，Channel  %s", msg.Header.Channel)
	return 0
}
