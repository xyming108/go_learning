package main

import (
	"fmt"
	"sync"
)

/**
Author: xym
Date: 2021/4/10 9:34
Project: Go_Learning
Description:
*/

//匿名函数的方法

var wg1 sync.WaitGroup

func main() { //开启一个主goroutine去执行main函数

	wg1.Add(10)
	for i := 0; i < 10; i++ {
		go func(m int) {
			fmt.Println("hello ", m)
			wg1.Done()
		}(i)
	}

	fmt.Println("hello main")

	wg1.Wait() //等所有小弟都干完活之后才收兵，即计数器为0时才结束
}
