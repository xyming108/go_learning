package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// 指定字段类型
func main() {
	var data = []byte(`{"status": 200}`)
	var result map[string]interface{}

	var decoder = json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()

	if err := decoder.Decode(&result); err != nil {
		log.Fatalln(err)
	}

	var status, _ = result["status"].(json.Number).Int64()
	fmt.Println("\nStatus value: ", status)
	fmt.Printf("%T\n", status)
}


// 你可以使用 string 来存储数值数据，在 decode 时再决定按 int 还是 float 使用
// 将数据转为 decode 为 string
/*func main() {
	var data = []byte(`{"status": 200}`)
	var result map[string]interface{}
	var decoder = json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	if err := decoder.Decode(&result); err != nil {
		log.Fatalln(err)
	}
	var status uint64
	err := json.Unmarshal([]byte(result["status"].(json.Number).String()), &status)
	checkError11(err)
	fmt.Println("\nStatus value: ", status)
	fmt.Printf("%T\n", status)
}*/

func checkError11(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
