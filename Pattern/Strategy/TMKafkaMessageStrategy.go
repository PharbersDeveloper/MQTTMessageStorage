package Strategy

import (
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/alfredyang1986/blackmirror/bmlog"
)

type TMKafkaMessageStrategy struct {
	Rd *BmRedis.BmRedis
	URI string
}

func (tkms *TMKafkaMessageStrategy) DoExecute(msg Message) (interface{}, error) {
	bmlog.StandardLogger().Info("TMKafkaMessageStrategy DoExecute")
	return nil, nil
}
