package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// 将 decode 的值转为 int 使用
func main() {
	var data = []byte(`{"status": 200}`)
	var result map[string]interface{}

	if err := json.Unmarshal(data, &result); err != nil {
		log.Fatalln(err)
	}

	var status = int(result["status"].(float64))
	fmt.Println("Status value: ", status)
	fmt.Printf("%T", status)
}
