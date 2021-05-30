package main

import "fmt"

//自定义类型和类型别名

//自定义类型
type myInt int

//类型别名
type yourInt = int

func main() {
	var n myInt
	n = 100
	fmt.Println(n)
	fmt.Printf("%T\n", n) //main.myInt

	var m yourInt
	m = 200
	fmt.Println(m)
	fmt.Printf("%T\n", m) //int
}
