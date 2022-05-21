package main

import (
	"fmt"
	"sync"
)

var g sync.WaitGroup

func syncMapThread(l int) {
	var ma sync.Map
	for i := 0; i < l; i++ {
		g.Add(1)
		go func(i int) {
			defer g.Done()
			ma.Store(i, i)
		}(i)
	}
	g.Wait()
	ma.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
}

func main() {
	syncMapThread(10)
}
