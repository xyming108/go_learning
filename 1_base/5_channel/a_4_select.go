package main

import "fmt"

/**
Author: xym
Date: 2021/4/10 19:55
Project: Go_Learning
Description:
*/

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		default:
			fmt.Println("什么都不干")
		}
	}
}
