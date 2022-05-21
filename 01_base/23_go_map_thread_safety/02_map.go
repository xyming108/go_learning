package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mm sync.Mutex
func mapThread(l int) {
	mp := make(map[int]int)
	for i := 0; i < l; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mm.Lock()
			mp[i] = i
			mm.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Println(len(mp))
}

func main() {
	mapThread(10)
}
