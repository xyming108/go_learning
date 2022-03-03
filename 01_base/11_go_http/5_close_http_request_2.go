package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
你可以创建一个自定义配置的 HTTP transport 客户端，用来取消 HTTP 全局的复用连接：

根据需求选择使用场景：
若你的程序要向同一服务器发大量请求，使用默认的保持长连接。
若你的程序要连接大量的服务器，且每台服务器只请求一两次，那收到请求后直接关闭连接。
或增加最大文件打开数 fs.file-max 的值。
*/

func main() {
	tr := http.Transport{DisableKeepAlives: true}
	client := http.Client{Transport: &tr}

	resp, err := client.Get("https://golang.google.cn/")
	if resp != nil {
		defer resp.Body.Close()
	}
	checkError5(err)

	fmt.Println(resp.StatusCode) // 200

	body, err := ioutil.ReadAll(resp.Body)
	checkError5(err)

	fmt.Println(len(string(body)))
}

func checkError5(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
