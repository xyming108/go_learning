package main

import "fmt"

//匿名函数
var s1 = func(x, y int) {
	fmt.Println(x + y)
}

func main() {
	s1(10, 20)

	//函数内部没有办法声明有名字的函数
	m1 := func(x, y int) {
		fmt.Println(x - y)
	}
	m1(20, 19)

	//如果只是调用一次的函数，还可以简写成立即执行函数，在后面多加一个括号
	func() {
		fmt.Println("Hello World")
	}()
	func(x, y int) {
		fmt.Println(x * y)
	}(100, 301)
}
