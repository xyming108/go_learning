package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
一些支持 HTTP1.1 或 HTTP1.0 配置了 connection: keep-alive 选项的服务器会保持一段时间的长连接。
但标准库 "net/http" 的连接默认只在服务器主动要求关闭时才断开，所以你的程序可能会消耗完 socket 描述符。
解决办法有 2 个，请求结束后：

直接设置请求变量的 Close 字段值为 true，每次请求结束后就会主动关闭连接。
设置 Header 请求头部选项 Connection: close，然后服务器返回的响应头部也会有这个选项，
此时 HTTP 标准库会主动断开连接。
*/

// 主动关闭连接
func main() {
	req, err := http.NewRequest("GET", "http://golang.or", nil)
	checkError3(err)

	req.Close = true
	//req.Header.Add("Connection", "close")    // 等效的关闭方式

	resp, err := http.DefaultClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	checkError3(err)

	body, err := ioutil.ReadAll(resp.Body)
	checkError3(err)

	fmt.Println(string(body))
}

func checkError3(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
