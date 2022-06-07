package main

import (
	"context"
	"fmt"
	"time"
)

func gen(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			select {
			case <-ctx.Done():
				fmt.Println("----------Done!!")
				return
			case ch <- n:
				fmt.Println("----------n:", n)
				n++
				time.Sleep(time.Millisecond * 200)
			}
		}
	}()
	return ch
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 避免其他地方忘记 cancel，且重复调用不影响

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			cancel()
			fmt.Println("--", ctx.Err())
			break
		}
	}
}
