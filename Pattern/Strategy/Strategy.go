package Strategy

import "MQTTStorage/Model"

type MessageStrategy interface {
	DoExecute(msg Model.Message) (interface{}, error)
}


