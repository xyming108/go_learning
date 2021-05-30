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
	runtime.GOMAXPROCS(6)
	wg2.Add(2)
	go a()
	go b()
	wg2.Wait()
}
