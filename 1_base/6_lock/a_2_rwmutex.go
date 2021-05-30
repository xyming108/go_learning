package main

import (
	"fmt"
	"sync"
	"time"
)

/**
Author: xym
Date: 2021/4/10 21:57
Project: Go_Learning
Description:
*/

var (
	s      int64
	wg1    sync.WaitGroup
	lock1  sync.Mutex
	rwLock sync.RWMutex
)

func read() {
	//lock1.Lock()
	rwLock.RLock()
	time.Sleep(time.Millisecond)
	rwLock.RUnlock()
	//lock1.Unlock()
	wg1.Done()
}

func write() {
	//lock1.Lock()
	rwLock.RLock()
	s = s + 1
	time.Sleep(time.Millisecond * 10)
	rwLock.RUnlock()
	//lock1.Unlock()
	wg1.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg1.Add(1)
		go read()
	}

	for i := 0; i < 10; i++ {
		wg1.Add(1)
		go write()
	}
	wg1.Wait()
	fmt.Println(time.Now().Sub(start))
}
