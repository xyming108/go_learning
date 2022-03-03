package main

import (
	"fmt"
	"time"
)

func main() {
	inCh := make(chan int)
	outCh := make(chan int)
	go func() {
		var in <-chan int = inCh
		var out chan<- int
		var val int
		for {
			select {
			case out <- val:
				println("--------", val)
				out = nil
				in = inCh
				time.Sleep(500*time.Millisecond)
				fmt.Println()
			case val = <-in:
				println("++++++++++", val)
				out = outCh
				in = nil
			}
		}
	}()
	go func() {
		for r := range outCh {
			fmt.Println("Result: ", r)
		}
	}()
	time.Sleep(0)
	inCh <- 1
	inCh <- 2
	time.Sleep(3 * time.Second)
}
