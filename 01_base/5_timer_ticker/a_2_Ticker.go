package main

import (
	"fmt"
	"time"
)

/*
Ticker：时间到了，多次执行
*/

func main() {
	// 1.获取ticker对象
	ticker := time.NewTicker(200 * time.Millisecond)
	i := 0
	// 子协程
	go func() {
		for {
			//<-ticker.C
			i++
			fmt.Println(<-ticker.C)
			fmt.Println(i)
			if i == 5 {
				//停止
				ticker.Stop()
			}

		}
	}()
	time.Sleep(time.Second*5)
}
