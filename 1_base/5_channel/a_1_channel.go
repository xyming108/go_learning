package main

import "fmt"

/**
Author: xym
Date: 2021/4/10 11:34
Project: Go_Learning
Description:
*/

// 声明： var 变量 channel 元素类型
// 初始化： make(chan 元素类型, [缓冲大小])
// 三种操作：send、receive、close
// 发送和接收都使用 <- 符号

func main() {
	var ch1 chan int        //引用类型，需要初始化
	ch1 = make(chan int, 1) //有缓冲通道
	ch1 <- 20               //发送到通道
	x := <-ch1              //从通道接收
	fmt.Println(x)
	close(ch1) //关闭通道
}
