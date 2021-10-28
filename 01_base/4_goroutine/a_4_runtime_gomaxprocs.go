package main

import (
	"fmt"
	"runtime"
	"sync"
)

/**
Author: xym
Date: 2021/4/10 9:52
Project: Go_Learning
Description:
*/

/*
Go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个OS线程来同时执行Go代码。
默认值是机器上的CPU核心数。例如在一个8核心的机器上，调度器会把Go代码同时调度到8个OS线程上（GOMAXPROCS是m:n调度中的n）。

	Go语言中的操作系统线程和goroutine的关系：

	1.一个操作系统线程对应用户态多个goroutine。
	2.go程序可以同时使用多个操作系统线程。
	3.goroutine和OS线程是多对多的关系，即m:n。

Go语言中可以通过runtime.GOMAXPROCS()函数设置当前程序并发时占用的CPU逻辑核心数。

Go1.5版本之前，默认使用的是单核心执行。Go1.5版本之后，默认使用全部的CPU逻辑核心数。
*/

var wg2 sync.WaitGroup

func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("A:", i)
	}
	wg2.Done()
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("B:", i)
	}
	wg2.Done()
}

func main() {
	runtime.GOMAXPROCS(2)
	wg2.Add(2)
	go a()
	go b()
	wg2.Wait()
}

/*

当runtime.GOMAXPROCS(1)设置为1时，只用单核CPU，此时是做完一个任务再做另一个任务
B: 0
B: 1
B: 2
B: 3
B: 4
B: 5
B: 6
B: 7
B: 8
B: 9
A: 0
A: 1
A: 2
A: 3
A: 4
A: 5
A: 6
A: 7
A: 8
A: 9
当runtime.GOMAXPROCS(2)设置为2时，用双核CPU，此时顺序可能会打乱
B: 0
B: 1
B: 2
B: 3
B: 4
B: 5
B: 6
A: 0
A: 1
A: 2
A: 3
A: 4
A: 5
A: 6
A: 7
A: 8
A: 9
B: 7
B: 8
B: 9

*/

