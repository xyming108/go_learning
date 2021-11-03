package main

import "fmt"

func main() {
	//var n = 5
	switch n := 5; n {
	//此时n只在该循环有效
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("无名指")
	default:
		fmt.Println("不存在")
	}
}
