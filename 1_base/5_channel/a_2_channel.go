package main

import "fmt"

/**
Author: xym
Date: 2021/4/10 11:52
Project: Go_Learning
Description:
*/

/*
	两个goroutine
		1、生成0-100的数字发送到ch2
		2、从ch2中取出数据计算它的平方，把结果发送到ch3中
*/

func f1(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

func f2(ch1 chan int, ch2 chan int) {
	for {
		tmp, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- tmp * tmp
	}
	close(ch2)
}

func main() {
	ch2 := make(chan int, 100)
	ch3 := make(chan int, 200)

	go f1(ch2)
	go f2(ch2, ch3)

	for ret := range ch3 {
		fmt.Println(ret)
	}
}
