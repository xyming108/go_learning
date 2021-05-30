package main

import "fmt"

/**
Author: 1_base
Date: 2021/4/6 11:45
Project: Go_Learning
Description:
*/

//结构体模拟实现其他语言中的继承

type animal struct {
	name string
}

//给animal实现一个移动的方法
func (a animal) move() {
	fmt.Printf("%s会动！", a.name)
}

//狗类
type dog2 struct {
	feet uint8
	animal
}

//给狗实现一个汪汪汪的方法
func (d dog2) wang() {
	fmt.Printf("%s在叫，汪汪汪~", d.name)
}

func main() {
	d1 := dog2{
		animal: animal{"咔哒"},
		feet:   4,
	}
	fmt.Println(d1)
	d1.wang()
	d1.move()
}
