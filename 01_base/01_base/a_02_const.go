package main

import "fmt"

//批量声明常量
const (
	statusOk = 200
	notFound = 404
)

const (
	//批量声明常量时，如果某一行声明后没有赋值，默认就和上一行一致
	a = 100
	b
	c
)

func main() {
	fmt.Println(statusOk)
	fmt.Println(notFound)
	fmt.Println(a, b, c)
}
