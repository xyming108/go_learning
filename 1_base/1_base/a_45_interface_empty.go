package main

import "fmt"

/**
Author: 1_base
Date: 2021/4/8 21:18
Project: Go_Learning
Description:
*/

//空接口
//所有的类型都实现了空接口，也就是任意类型的变量都能保存到空接口中
//interface：关键字
//interface{}：空接口类型

func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func main() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "周林"
	m1["age"] = 9000
	m1["married"] = true
	m1["hobby"] = [...]string{"唱", "跳", "rap"}
	fmt.Println(m1)

	show(false)
	show(nil)
	show(m1)
}
