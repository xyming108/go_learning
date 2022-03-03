package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
在 encode/decode JSON 数据时，Go 默认会将数值当做 float64 处理
*/

func main() {
	var data = []byte(`{"status": 200}`)
	var result map[string]interface{}

	if err := json.Unmarshal(data, &result); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("\n%T\n\n", result["status"]) // float64

	var status = result["status"].(float64) // 类型断言正确
	fmt.Println("Status value: ", status)

	var statuss = result["status"].(int) // 类型断言错误
	fmt.Println("Status value: ", statuss)
}
