package main

import "fmt"

/**
Author: 01_base
Date: 2021/4/8 18:15
Project: Go_Learning
Description:
*/

//同一个结构体可以实现多个接口
//接口还可以嵌套

type animal5 interface {
	mover
	eater
}

type mover interface {
	move1()
}

type eater interface {
	eat1(string)
}

type cat5 struct {
	name string
	feet int8
}

//cat5实现了mover接口和eater接口
func (c *cat5) move1() {
	fmt.Println("走猫步！")
}

func (c *cat5) eat1(food string) {
	fmt.Println("猫吃鱼", food)
}

func main() {
	var m mover
	var e eater
	var aa animal5
	c5 := &cat5{
		name: "暗示法",
		feet: 100,
	}
	aa = c5
	aa.move1()
	aa.eat1("你你你")
	fmt.Println()

	m = c5
	m.move1()
	fmt.Println(m)
	fmt.Println()

	e = c5
	e.eat1("哈哈")
	fmt.Println(e)
}
