package main

import "fmt"

/**
Author: 1_base
Date: 2021/4/7 14:08
Project: Go_Learning
Description:
*/

//接口：是一种特殊的类型，它规定了变量有哪些方法

type car interface {
	run()
}

type falali struct {
	brand string
}

type baoshijie struct {
	brand string
}

func (f falali) run() {
	fmt.Printf("%s速度70迈", f.brand)
}

func (b baoshijie) run() {
	fmt.Printf("%s速度100迈", b.brand)
}

func drive(c car) {
	c.run()
}

func main() {
	var f1 = falali{
		brand: "法拉利",
	}
	var b1 = baoshijie{
		brand: "保时捷",
	}
	drive(f1)
	fmt.Println()
	drive(b1)
}
