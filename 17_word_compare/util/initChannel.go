package util

type MsgStruct struct {
	Sn   string   `json:"sn"`
	Urls []string `json:"urls"`
}

var ch chan MsgStruct

func InitChannel() {
	ch = make(chan MsgStruct, 20)
}

func GetChan() chan MsgStruct {
	return ch
}

func SendToChan(msg MsgStruct) {
	ch <- msg
}

type LarkMsgStruct struct {
	Image       string `json:"image"`
	Sku         string `json:"sku"`
	Position    string `json:"position"`
	ReSendCount int    `json:"re_send_count"` //重发次数
}

const MaxLimit = 20         // 最大请求数
const PerMillisecond = 1000 // 间隔请求时间 ms

var larkCh chan LarkMsgStruct

func InitLarkChannel() {
	larkCh = make(chan LarkMsgStruct, MaxLimit)
}

func GetLarkChan() chan LarkMsgStruct {
	return larkCh
}

func SendToLarkChan(msg LarkMsgStruct) {
	larkCh <- msg
}
