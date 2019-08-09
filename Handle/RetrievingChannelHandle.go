package Handle

import (
	"encoding/json"
	"github.com/PharbersDeveloper/MQTTMessageStorage/Daemons"
	"github.com/PharbersDeveloper/MQTTMessageStorage/Model"
	"github.com/PharbersDeveloper/MQTTMessageStorage/Pattern/Strategy"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/alfredyang1986/blackmirror/bmlog"
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

func (rc RetrievingChannelHandler) RetrievingChannel(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	w.Header().Add("Content-Type", "application/json")
	var response map[string]interface{}
	response = make(map[string]interface{})

	enc := json.NewEncoder(w)
	body, err := ioutil.ReadAll(r.Body)
	msg := Model.Message{}
	err = json.Unmarshal(body, &msg)

	if err != nil {bmlog.StandardLogger().Error(err); return 1}
	context := Strategy.MessageContext{ Msg: msg, Rd: rc.rd, Em: rc.em }
	_, err = context.DoExecute()
	if err != nil {
		response["status"] = "error"
		response["code"] = 500
		response["msg"] = err.Error()
		enc.Encode(response)
		return 1
	}
	return 0
}
