package main

import (
	"fmt"
	"math/rand"
	"time"
)

var ch chan int

func init() {
	ch = make(chan int, 10)
}

func main() {
	tk := time.NewTicker(time.Second)
	defer tk.Stop()
	go func() {
		for {
			select {
			case <-tk.C:
				r := rand.Intn(10)
				fmt.Println("=====r: ", r)
				ch <- r
			}
		}
	}()

	go func() {
		for {
			select {
			case r, ok := <-ch:
				if ok {
					fmt.Println("+++++r: ", r)
					time.Sleep(time.Second * 2)
				}
			}
		}
	}()

	for {

	}
}
