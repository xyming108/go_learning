//++++++++++++++++++++++++++++++++++++++++
// 《Go语言高级开发与实战》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 知乎：https://www.zhihu.com/people/shirdonl
// 公众号:源码大数据
// 仓库地址：https://gitee.com/shirdonl/goAdvanced
// 仓库地址：https://github.com/shirdonl/goAdvanced
//++++++++++++++++++++++++++++++++++++++++

package util

import (
	"encoding/json"
	"log"
	"net/http"
)

//返回数据结构体
type ResponseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

//返回json格式的数据给客户端
func ResponseJson(writer http.ResponseWriter, code int, msg string, data interface{}) {
	Response(writer, code, msg, data)
}

func Response(writer http.ResponseWriter, code int, msg string, data interface{}) {
	//设置header为JSON 默认是test/html，所以特别指出返回的数据类型为application/json
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	rep := ResponseData{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	//将结构体转化为json字符串
	ret, err := json.Marshal(rep)
	if err != nil {
		log.Panicln(err.Error())
	}

	//返回json ok
	writer.Write(ret)
}
