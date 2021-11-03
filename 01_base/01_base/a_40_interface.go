package main

import "fmt"

/**
Author: 01_base
Date: 2021/4/7 14:08
Project: Go_Learning
Description:
*/

//接口：是一种特殊的类型，它规定了变量有哪些方法

type speaker interface {
	speak()
}

type cat struct {
}

type fish struct {
}

type duck struct {
}

func (c cat) speak() {
	fmt.Println("喵喵喵~")
}

func (f fish) speak() {
	fmt.Println("咕咕咕~")
}

func (d duck) speak() {
	fmt.Println("汪汪汪~")
}

func aa(speaker2 speaker) {
	speaker2.speak()
}

func main() {
	var c1 cat
	var f1 fish
	var d1 duck

	aa(c1)
	aa(f1)
	aa(d1)

	var ss speaker
	ss = c1
	fmt.Println(ss)
}
