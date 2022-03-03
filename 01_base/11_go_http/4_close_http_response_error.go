package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// 请求失败造成 panic
func main() {
	resp, err := http.Get("https://www.sulvblog")
	defer resp.Body.Close()    // resp 可能为 nil，不能读取 Body
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	fmt.Println(string(body))
}

func checkError(err error) {
	if err != nil{
		log.Fatalln(err)
	}
}