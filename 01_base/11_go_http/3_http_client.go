package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

//客户端
func client() {
	resp, err := http.Get("http://127.0.0.1:8000/go")
	if err != nil {
		log.Fatal("url请求错误")
	}
	defer resp.Body.Close()
	buf := make([]byte, 1024)
	for {
		read, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("err: ", err)
			return
		} else {
			fmt.Println("读取完毕")
			res := string(buf[:read])
			fmt.Println("result: ", res)
			break
		}
	}
}

func main() {
	client()
}
