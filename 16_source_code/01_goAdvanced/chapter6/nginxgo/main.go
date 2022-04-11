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

package main

import (
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/secKill/goods/1", handleController)
	http.ListenAndServe(":8081", nil)
}

//处理请求函数，根据请求将响应结果信息写入日志
func handleController(w http.ResponseWriter, r *http.Request) {
	failedMsg := "handle in port:"
	writeLog(failedMsg, "./request.log")
}

//写入日志
func writeLog(msg string, logPath string) {
	fd, _ := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer fd.Close()
	content := strings.Join([]string{msg, "\r\n"}, "8081")
	buf := []byte(content)
	fd.Write(buf)
}
