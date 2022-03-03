package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

/*
使用 struct 类型将你需要的数据映射为数值型
*/

// struct 中指定字段类型
func main() {
	var data = []byte(`{"status": 200}`)
	var result struct {
		Status uint64 `json:"status"`
	}

	err := json.NewDecoder(bytes.NewReader(data)).Decode(&result)
	checkError111(err)
	fmt.Printf("\nResult: %+v", result)
	fmt.Printf("\n%T", result.Status)
}

func checkError111(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
