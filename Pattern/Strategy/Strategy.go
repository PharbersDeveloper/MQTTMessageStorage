package Strategy

import "github.com/PharbersDeveloper/MQTTMessageStorage/Model"

type MessageStrategy interface {
	DoExecute(msg Model.Message) (interface{}, error)
}


