package main

import (
	"fmt"
)

/**
Author: 01_base
Date: 2021/4/5 22:47
Project: Go_Learning
Description:
*/

//给自定义类型加方法
//不能给别的包里面添加方法，只能给自己包里的类型添加方法

type myInt1 int

func (m myInt1) hello() {
	fmt.Println("我是一个int", m)
}

//结构体初始化
type person3 struct {
	name string
	age  int
}

//为什么要有构造函数
func newPerson3(x string, y int) person3 {
	return person3{
		name: x,
		age:  y,
	}
}

/*func newPerson3(x string, y int) *person3 {
	return &person3{
		name: x,
		age:  y,
	}
}*/

func main() {
	//声明一个int32类型的变量x，他的值是10
	//方法1：
	var x1 int32
	x1 = 10
	//方法2：
	var x2 int32 = 10
	//方法3：
	var x3 = int32(10)
	//方法4：
	x4 := 10
	fmt.Println(x1, x2, x3, x4)

	//声明一个myInt1类型的变量m，它的值是100
	//方法1：
	var m1 myInt1
	m1 = 100
	//方法2：
	var m2 myInt1 = 100
	//方法3：
	var m3 = myInt1(100)
	//方法4：
	m4 := myInt1(100)

	m1.hello()
	m2.hello()
	m3.hello()
	m4.hello()

	//方法1：
	var p person3
	p.name = "元帅"
	p.age = 18
	fmt.Println(p)
	var p2 person3
	p2.name = "元帅"
	p2.age = 18
	fmt.Println(p2)
	//方法2：
	var pp1 = person3{
		name: "沙姜粉",
		age:  10,
	}
	fmt.Println(pp1)
}
