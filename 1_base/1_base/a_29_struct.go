package main

import (
	"fmt"
)

//结构体

type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	var p person
	p.name = "信息员"
	p.age = 10000
	p.gender = "男"
	p.hobby = []string{"乒乓球", "写代码"}
	fmt.Println(p)
	fmt.Println()

	var pp = new(person)
	(*pp).name = "信息员"
	pp.age = 1000
	pp.gender = "女"
	pp.hobby = []string{"乒乓", "代码"}
	fmt.Println(pp)
	fmt.Printf("%T\n", pp)
	fmt.Println(*pp)
	fmt.Println()

	//若没加&，是对结构体进行初始化
	//若加了&，则是对结构体指针进行初始化
	var ppp = &person{
		name:   "信息员",
		gender: "男",
	}
	//使用值列表的形式初始化，值的顺序要和结构体定义时字段的顺序一致
	var poppy = person{
		"信息",
		8,
		"女",
		[]string{"乒乓球", "写代码"},
	}
	fmt.Println(ppp, "\n", poppy)

	//匿名结构体，多用于一些临时场景
	var s struct {
		name string
		age  int
	}
	s.name = "嘿嘿嘿"
	s.age = 100
	fmt.Println(s)
}
