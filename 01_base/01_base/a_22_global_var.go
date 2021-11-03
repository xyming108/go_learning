package main

import "fmt"

//变量的作用域

//定义一个全局变量
var x = 100

func p1() {
	fmt.Println(x)
}

func main() {
	p1()
}
