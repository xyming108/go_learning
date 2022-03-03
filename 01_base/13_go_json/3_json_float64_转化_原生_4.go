package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

/*
可以使用 struct 将数值类型映射为 json.RawMessage 原生数据类型
适用于如果 JSON 数据不着急 decode 或 JSON 某个字段的值类型不固定等情况：
*/

// 状态名称可能是 int 也可能是 string，指定为 json.RawMessage 类型
func main() {
	records := [][]byte{
		[]byte(`{"status":200, "tag":"one"}`),
		[]byte(`{"status":"ok", "tag":"two"}`),
	}

	for idx, record := range records {
		var result struct {
			StatusCode uint64
			StatusName string
			Status     json.RawMessage `json:"status"`
			Tag        string          `json:"tag"`
		}

		err := json.NewDecoder(bytes.NewReader(record)).Decode(&result)
		checkError1111(err)

		var name string
		err = json.Unmarshal(result.Status, &name)
		if err == nil {
			result.StatusName = name
		}

		var code uint64
		err = json.Unmarshal(result.Status, &code)
		if err == nil {
			result.StatusCode = code
		}

		fmt.Printf("\n[%v] result => %+v\n", idx, result)
	}
}

func checkError1111(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
