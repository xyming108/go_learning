package main

import (
	"fmt"
)

func main() {
	/*go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")
	// 主协程
	for i := 0; i < 2; i++ {
		// 切一下，再次分配任务
		runtime.Gosched() //让出时间片
		fmt.Println("hello")
	}*/
	/*
		结果总是：
		world
		world
		hello
		hello
	*/

	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			// 结束协程
			//runtime.Goexit() //退出当前协成
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	for{
	}

	/*
		执行到：
		B.defer
		A.defer
		就不会往下走了
	*/
}
