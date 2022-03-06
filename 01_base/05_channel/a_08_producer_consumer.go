package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func producer(data chan<- int) {
	for i := 0; i < 4; i++ {
		data <- i
	}
	close(data)
}

func consumer(data <-chan int) {
	defer wg.Done()
	for {
		v, ok := <-data
		if !ok {
			break
		}
		fmt.Println("---1:", v, "===1:", ok)
	}
}

func main() {
	data := make(chan int)
	go producer(data)

	wg.Add(1)
	go consumer(data)
	wg.Wait()
}
