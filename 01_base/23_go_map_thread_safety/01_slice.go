package main

import (
	"fmt"
	"sync"
)

var m sync.Mutex
var w sync.WaitGroup

func sliceSafety(l int) {
	var s []int
	var sum int
	fmt.Printf("----------: len(s): %d, cap(s): %d, s: %v \n", len(s), cap(s), s)
	for i := 0; i < l; i++ {
		w.Add(1)
		go func(i int) {
			defer w.Done()
			m.Lock()
			sum++
			s = append(s, i)
			fmt.Printf("==========i: %d: len(s): %d, cap(s): %d, s: %v \n", i, len(s), cap(s), s)
			m.Unlock()
		}(i)
	}
	w.Wait()
	fmt.Println(sum)
	fmt.Println(s, len(s))
}

func main() {
	sliceSafety(10)
}
