package Strategy

type MessageStrategy interface {
	DoExecute(msg Message) (interface{}, error)
}


