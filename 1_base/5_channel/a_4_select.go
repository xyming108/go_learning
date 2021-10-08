package main

import (
	"fmt"
	"time"
)

/**
Author: xym
Date: 2021/4/10 19:55
Project: Go_Learning
Description:
*/

func main() {
	/*ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		default:
			fmt.Println("什么都不干")
		}
	}*/

	//-------------------------------------

	/*// 创建2个管道
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	go func() {
		//time.Sleep(2 * time.Second)
		int_chan <- 1
	}()
	go func() {
		string_chan <- "hello"
	}()
	//如果多个channel同时ready，则随机选择一个执行
	select {
	case value := <-int_chan:
		fmt.Println("int:", value)
	case value := <-string_chan:
		fmt.Println("string:", value)
	}
	fmt.Println("main结束")*/

	//---------------------------------------
	// 2个管道
	output1 := make(chan string)
	output2 := make(chan string)
	// 跑2个子协程，写数据
	go func(ch chan string) {
		time.Sleep(time.Second * 1)
		ch <- "test1"
	}(output1)
	go func(ch chan string) {
		time.Sleep(time.Second * 2)
		ch <- "test2"
	}(output2)
	// select可以同时监听一个或多个channel，直到其中一个channel ready
	select {
	case s1 := <-output1:
		fmt.Println("s1=", s1)
	case s2 := <-output2:
		fmt.Println("s2=", s2)
	}
}
