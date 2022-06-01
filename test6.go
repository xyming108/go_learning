package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1: ", <-ch)
	}()
	fmt.Println("发送前")
	ch <- 123
	fmt.Println("发送后")
	fmt.Println("第二次发送前")
	ch <- 456
	fmt.Println("第二次发送后")
	wg.Wait()
}
