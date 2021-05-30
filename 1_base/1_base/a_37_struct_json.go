package main

import (
	"encoding/json"
	"fmt"
)

/**
Author: 1_base
Date: 2021/4/6 12:03
Project: Go_Learning
Description:
*/

//结构体与json

//1、序列化：把Go语言中的结构体变量 --> json格式的字符串
//2、反序列化：json格式的字符串 --> Go语言中能够识别的结构体变量

type person7 struct {
	Name string //首字母大写
	Age  int
}

func main() {
	p := person7{
		Name: "利可君",
		Age:  18,
	}
	//序列化
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("marshal failed, err: %v", err)
		return
	}
	fmt.Printf("%#v\n", string(b))
	//反序列化
	str := `{"name":"理想", "age":18}`
	var pp person7
	json.Unmarshal([]byte(str), &pp) //指针是为了能在json.Unmarshal内部修改pp的值
	fmt.Printf("%#v\n", pp)
}
