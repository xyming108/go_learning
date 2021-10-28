package main

import "fmt"

/**
Author: 01_base
Date: 2021/4/6 11:02
Project: Go_Learning
Description:
*/

//匿名字段
//不常用

type person5 struct {
	string
	int
}

func main() {
	p1 := person5{
		"嗯嗯嗯",
		900,
	}
	fmt.Println(p1)
	fmt.Println(p1.string)
	fmt.Println(p1.int)
}
