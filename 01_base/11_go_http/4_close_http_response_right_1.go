package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// 大多数情况正确的示例
func main() {
	//应该先检查 HTTP 响应错误为 nil，再调用 resp.Body.Close() 来关闭响应体
	resp, err := http.Get("https://www.sulvblog.c")
	checkErrorr(err)

	defer resp.Body.Close() // 绝大多数情况下的正确关闭方式
	body, err := ioutil.ReadAll(resp.Body)
	checkErrorr(err)

	fmt.Println(string(body))
}

func checkErrorr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
