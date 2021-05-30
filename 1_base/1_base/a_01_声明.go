package main

import "fmt"

//声明变量
var name string
var age int
var isOk bool

//批量声明
var (
	a1 string // ""
	a2 int    // 0
	a3 bool   // false
)

func main() {
	name = "理想"
	age = 18
	isOk = true

	a1 = "啊啊"
	a2 = 20
	a3 = true

	//Go语言中非全局变量声明后必须要使用，不使用就无法通过编译
	var pp int
	fmt.Println(name, age, isOk, a1, a2, a3, pp)

	//声明变量同时赋值
	var s1 string = "xyMing"
	fmt.Println(s1)
	//类型推导（自动）
	var s2 = "xyMing"
	fmt.Println(s2)
	//简短变量声明（只能在函数中用，即不能在全局使用）
	s3 := "哈哈哈"
	fmt.Println(s3)

	//匿名变量（一个下划线表示：_）

}
