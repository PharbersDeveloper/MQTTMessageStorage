package Model
type header struct {
	Method 		string	`json:"method"`
	Channel		string	`json:"channel"`
	Topic		string 	`json:"topic"`
}

type Message struct {
	Header 	header			`json:"header"`

	// PayLoad具体不需要知道，因每个Interface和服务都不一样，作为转发的中间件其实不需要理会，只需要知道中间件提供的接口参数的格式即可
	PayLoad	interface{}		`json:"payload"`
}