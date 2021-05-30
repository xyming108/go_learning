package main

import (
	"fmt"
	"sync"
)

/**
Author: xym
Date: 2021/4/9 22:36
Project: Go_Learning
Description:
*/

//一个操作系统线程对应用户态多个goroutine
//go程序可以同时使用多个操作系统线程
//goroutine和OS线程是多对多的关系，即m:n

var wg sync.WaitGroup

func hello(i int) {
	fmt.Println("hello ", i)
	wg.Done() //通知wg，表示已经完成这个函数，计数器-1
}

func main() { //开启一个主goroutine去执行main函数
	/*wg.Add(10)*/ //也可以直接使用这种一次性加满
	for i := 0; i < 10; i++ {
		wg.Add(1)   //计数牌+1
		go hello(i) //开启了一个goroutine去执行hello这个函数
	}

	fmt.Println("hello main")
	/*time.Sleep(time.Second)*/ //这个不能随便用，因为它表示cpu真的会停留这么长时间

	wg.Wait() //等所有小弟都干完活之后才收兵，即计数器为0时才结束
}
