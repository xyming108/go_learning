package main

import (
	"fmt"
)

/**
Author: 01_base
Date: 2021/4/8 21:57
Project: Go_Learning
Description:
*/

//类型断言1
func assert(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Println("猜错了")
	} else {
		fmt.Println("传进来的是一个字符串", str)
	}
}

//类型断言2
func assert1(a interface{}) {
	fmt.Printf("%T\n", a)
	switch t := a.(type) {
	case string:
		fmt.Println("是一个字符串：", t)
	case int:
		fmt.Println("是一个int：", t)
	case int64:
		fmt.Println("是一个int64：", t)
	case bool:
		fmt.Println("是一个bool：", t)
	}
}

func main() {
	assert("字符串")
	fmt.Println()
	assert1("哈哈哈")
	fmt.Println()
	assert1(int64(200))
}
