package types

type BaseResponse struct {
	ResCode int    `json:"res_code" description:"1成功 0失败"`
	Message string `json:"message" description:"提示消息"`
}
