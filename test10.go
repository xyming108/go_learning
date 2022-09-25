package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(ii int) {
			defer wg.Done()

			defer func() {
				err := recover()
				fmt.Println("---------:", err)
			}()

			for j := ii; j < (ii+1)*1000; j++ { // 给goroutine划分执行区间
				if j == 5 {
					panic("发生了panic: ", ii)
				}
			}
		}(i)
	}
	wg.Wait()
}
