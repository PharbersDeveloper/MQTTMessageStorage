package Strategy

import (
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/alfredyang1986/blackmirror/bmlog"
)

type MAXKafkaMessageStrategy struct {
	Rd *BmRedis.BmRedis
	URI string
}

func (msms *MAXKafkaMessageStrategy) DoExecute(msg Message) (interface{}, error) {
	bmlog.StandardLogger().Info("MAXKafkaMessageStrategy DoExecute")
	return nil, nil
}