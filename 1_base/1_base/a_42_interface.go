package main

import (
	"fmt"
)

/**
Author: 1_base
Date: 2021/4/8 15:53
Project: Go_Learning
Description:
*/

type animal2 interface {
	move()
	eat(string)
}

type cat3 struct {
	name string
	feet int8
}

func (c cat3) move() {
	fmt.Println("走猫步！")
}

func (c cat3) eat(food string) {
	fmt.Println("猫吃鱼", food)
}

type chicken3 struct {
	feet int8
}

func (c chicken3) move() {
	fmt.Println("鸡动！")
}

func (c chicken3) eat(a string) {
	fmt.Println("吃鸡饲料！", a)
}

func main() {
	var a1 animal2
	bc := cat3{
		name: "淘气",
		feet: 4,
	}
	a1 = bc
	a1.eat("小红花")
	a1.move()
	fmt.Println(a1)
	fmt.Println()

	bc1 := chicken3{
		feet: 10,
	}
	a1 = bc1
	a1.move()
	a1.eat("哈哈哈")
	fmt.Println(a1)
}
