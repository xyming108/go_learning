package main

import "fmt"

/**
Author: xym
Date: 2021/4/10 11:34
Project: Go_Learning
Description:
*/

// 声明： var 变量 channel 元素类型
// 初始化： make(chan 元素类型, 缓冲大小)
// 三种操作：send、receive、close
// 发送和接收都使用 <- 符号
// 遵循 先进先出 的原则

func main() {
	var ch1 chan int        //引用类型，需要初始化
	ch1 = make(chan int, 1) //有缓冲通道
	//另一种初始化方法
	//ch1 := make(chan int, 1)
	ch1 <- 20  //发送到通道
	x := <-ch1 //从通道接收
	// <-ch1      // 从ch1中接收值，忽略结果
	fmt.Println(x)
	//只有在通知接收方goroutine所有的数据都发送完毕的时候才需要关闭通道，通道是可以被垃圾回收机制回收的，它和关闭文件是不一样的，在结束操作之后关闭文件是必须要做的，但关闭通道不是必须的。
	close(ch1) //关闭通道
	/*
		关闭后的通道有以下特点：
		    1.对一个关闭的通道再发送值就会导致panic。
		    2.对一个关闭的通道进行接收会一直获取值直到通道为空。
		    3.对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
		    4.关闭一个已经关闭的通道会导致panic。
	*/

	fmt.Println("----------------------------------")

	//无缓冲通道，有发送必须要有接收
	ch := make(chan int)
	go func(ch chan int) {
		ret := <-ch
		fmt.Println("接收成功", ret)
	}(ch) // 启用goroutine从通道接收值
	ch <- 10
	fmt.Println("发送成功")
	close(ch)
	/*
		无缓冲通道上的发送操作会阻塞，直到另一个goroutine在该通道上执行接收操作，这时值才能发送成功，两个goroutine将继续执行。
		相反，如果接收操作先执行，接收方的goroutine将阻塞，直到另一个goroutine在该通道上发送一个值。
	*/
}
