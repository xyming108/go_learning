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
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

var m map[string]string //缓存key-value
func main() {
	m = make(map[string]string, 10)
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler) // 将客户端发来的请求交给HomeHandler处理
	fmt.Println("listen...")
	http.ListenAndServe(":8080", r)
}
func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	// 解析客户端发来的命令
	argStr := r.RequestURI
	argStartIndex := strings.LastIndex(argStr, "?") + 1
	args := strings.Split(argStr[argStartIndex:], "|")
	for _, arg := range args {
		fmt.Println(arg)
	}
	command := args[0]
	// 处理命令
	switch command {
	case "get":
		fmt.Fprintf(rw, "%s = %s\n", args[1], handleGet(args[1]))
	case "set":
		handleSet(args[1], args[2])
		fmt.Fprintf(rw, "%s = %s\n", args[1], args[2])
	default:
		fmt.Println("command error")
	}
}
func handleGet(key string) string {
	if val, ok := m[key]; ok {
		return val
	}
	return ""
}
func handleSet(key, val string) {
	m[key] = val
}
