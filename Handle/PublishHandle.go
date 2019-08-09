package Handle

import (
	"encoding/json"
	"fmt"
	"github.com/PharbersDeveloper/MQTTMessageStorage/Daemons"
	"github.com/PharbersDeveloper/MQTTMessageStorage/Model"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/alfredyang1986/blackmirror/bmlog"
	emitter "github.com/emitter-io/go/v2"
	"github.com/go-redis/redis"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"reflect"
)

type PublishHandler struct {
	Method     	string
	HttpMethod 	string
	Args       	[]string
	rd         	*BmRedis.BmRedis
	em			*Daemons.Emitter
}

func (k PublishHandler) NewPublishHandler(args ...interface{}) PublishHandler {
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

	return PublishHandler{Method: md, HttpMethod: hm, Args: ag, rd: r, em: em}
}

func (k PublishHandler) GetHttpMethod() string {
	return k.HttpMethod
}

func (k PublishHandler) GetHandlerMethod() string {
	return k.Method
}

func (k PublishHandler) Publish(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	w.Header().Add("Content-Type", "application/json")
	var err error

	var response map[string]interface{}
	response = make(map[string]interface{})
	enc := json.NewEncoder(w)

	body, err := ioutil.ReadAll(r.Body)
	msg := Model.Message{}
	err = json.Unmarshal(body, &msg)
	if err != nil {bmlog.StandardLogger().Error(err); return 1}

	channel := msg.PayLoad.(map[string]interface{})["channel"].(string)
	rdClient := k.rd.GetRedisClient()
	result, err := rdClient.Get(fmt.Sprint("mqtt_channel_key_", channel)).Result()
	if err != redis.Nil {
		b, _ := json.Marshal(msg.PayLoad)
		client := k.em.GetClient()
		err = client.Publish(result, channel, b, emitter.WithAtLeastOnce())
		if err != nil { panic(err.Error()) }
	} else { panic(err.Error()) }

	response["status"] = "success"
	response["code"] = http.StatusOK
	response["msg"] = "Publish Success"
	enc.Encode(response)

	return 0
}
