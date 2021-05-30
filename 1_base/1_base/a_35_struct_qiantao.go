package main

import "fmt"

/**
Author: 1_base
Date: 2021/4/6 11:05
Project: Go_Learning
Description:
*/

//嵌套结构体

type address struct {
	province string
	city     string
}

type person6 struct {
	name    string
	age     int
	address //匿名嵌套结构体
	//workPlace
}

type company struct {
	name string
	addr address
}

type workPlace struct {
	province string
	city     string
}

func main() {
	p := person6{
		name: "萨克金凤凰",
		age:  50,
		address: address{
			"广东",
			"深圳",
		},
	}
	fmt.Println(p)
	fmt.Println(p.address.city)
	//如果结构体内只嵌套一个结构体，可用如下方法，否则只能用上面的方法
	fmt.Println(p.city)
}
