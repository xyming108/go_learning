package main

import (
	"fmt"
	"sync"
)

/**
Author: xym
Date: 2021/4/10 21:45
Project: Go_Learning
Description:
*/

//互斥锁是完全互斥的

var (
	x    int64
	wg   sync.WaitGroup
	lock sync.Mutex //互斥锁
)

func add() {
	for i := 0; i < 10000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
