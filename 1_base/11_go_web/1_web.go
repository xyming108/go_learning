package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

/**
Author: xym
Date: 2021/5/16 13:13
Project: myproject2_go_learning
Description:
*/

import (
	"log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	//这个写入到w的是输出到客户端的，在浏览器可以看见
	io.WriteString(w, "Hello from a HandleFunc #1!\n")
	fmt.Fprintf(w, "Hello astaxie!")

	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	/*
		输入：http://localhost:9090/?url_long=111&url_long=222
		浏览器得到：
			Hello from a HandleFunc #1!
			Hello astaxie!
		后台窗口得到：
			map[url_long:[111 222]]
			path /
			scheme
			[111 222]
			key: url_long
			val: 111222
			map[]
			path /favicon.ico
			scheme
			[]
	*/
}

func main() {
	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
