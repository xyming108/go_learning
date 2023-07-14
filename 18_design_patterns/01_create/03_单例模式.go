package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once
var singleInstance *single

type single struct{}

func GetSingleInstance() *single {
	if singleInstance == nil {
		once.Do(func() { // 只会执行一次，即使在并发的情况下
			fmt.Println("Create single instance.")
			singleInstance = &single{}
		})
	} else {
		fmt.Println("Get single instance.")
	}
	return singleInstance
}

func main() {
	for i := 0; i < 10; i++ {
		go GetSingleInstance()
	}

	time.Sleep(time.Second * 60)
}
