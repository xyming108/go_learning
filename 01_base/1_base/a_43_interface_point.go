package main

import "fmt"

/**
Author: 01_base
Date: 2021/4/8 17:52
Project: Go_Learning
Description:
*/

/*
使用值接收者和指针接收者的区别：
	使用值接收者实现接口，结构体类型和结构体指针类型的变量都能存
	指针接收者实现接口只能存结构体指针类型的变量
*/

/*
接口和类型的关系：
	多个类型可以实现同一个接口
	一个类型可以实现多个接口
*/

type animal4 interface {
	move()
	eat(string)
}

type cat4 struct {
	name string
	feet int8
}

//使用值接收者实现了接口的所有方法
/*
func (c cat4) move() {
	fmt.Println("走猫步！")
}

func (c cat4) eat(food string) {
	fmt.Println("猫吃鱼", food)
}
*/

//使用指针接收者实现了接口的所有方法
func (c *cat4) move() {
	fmt.Println("走猫步！")
}

func (c *cat4) eat(food string) {
	fmt.Println("猫吃鱼", food)
}

func main() {
	var a1 animal4

	c1 := cat4{"tom", 4}    //cat
	c2 := &cat4{"kevin", 5} //*cat

	/*
		a1 = c1
		fmt.Println(a1)
		a1 = c2
		fmt.Println(a1)
	*/

	a1 = &c1
	fmt.Println(a1)
	a1 = c2
	fmt.Println(a1)
}
