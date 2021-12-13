package main

import (
	"fmt"
	"log"
	"net/http"
)

//服务端
func server() {
	http.HandleFunc("/go", func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("连接成功: %+v", request.RemoteAddr)
		fmt.Println("url:", request.URL)
		fmt.Println("header:", request.Header)
		fmt.Println("body:", request.Body)

		writer.Write([]byte("我是服务端"))
	})
	err := http.ListenAndServe("127.0.0.1:8000", nil)
	if err != nil {
		return
	}
}

func main() {
	server()
}
