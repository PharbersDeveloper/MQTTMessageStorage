package MQTTChannelState

// 不要吐槽我，不想写channel，后续集群是用zk管理
var channels map[string]bool

type StateSlice struct {
	count int
}

func (ss *StateSlice) NewStateSlice() {
	channels = make(map[string]bool)
}

func (ss *StateSlice) Push(channelName string) {
	channels[channelName] = true
	ss.count = len(channels)
}

func (ss *StateSlice) GetCount() int {
	return ss.count
}

func (ss *StateSlice) Exist(channelName string) bool {
	_, b := channels[channelName]
	return b
}

