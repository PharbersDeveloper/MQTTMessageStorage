package Handler

import (
	"encoding/json"
	"github.com/PharbersDeveloper/MQTTMessageStorage/Pattern/Strategy"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/alfredyang1986/blackmirror/bmlog"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"reflect"
)

type KenGenHandler struct {
	Method     string
	HttpMethod string
	Args       []string
	rd         *BmRedis.BmRedis
}

func (k KenGenHandler) NewKenGenHandler(args ...interface{}) KenGenHandler {
	var r *BmRedis.BmRedis
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

	return KenGenHandler{Method: md, HttpMethod: hm, Args: ag, rd: r}
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

	body, err := ioutil.ReadAll(r.Body)
	msg := Strategy.Message{}
	err = json.Unmarshal(body, &msg)
	if err != nil {bmlog.StandardLogger().Error(err); return 1}

	context := Strategy.MessageContext{ Msg: msg, Rd: k.rd, URI: k.Args[0] }
	res, err := context.DoExecute()
	if err != nil { panic(err.Error()) }

	response["status"] = "success"
	response["code"] = 200
	response["msg"] = "Generate KeyGen Success"
	response["body"] = map[string]interface{}{ "channelKey": res.(string) }
	enc.Encode(response)

	return 0
}
