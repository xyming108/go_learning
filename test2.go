package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

var (
	stop  int32
	count int64
	sum   time.Duration
)

func concat() {
	for n := 0; n < 100; n++ {
		for i := 0; i < 8; i++ {
			go func() {
				s := "Go GC"
				s += " " + "Hello"
				s += " " + "World"
				_ = s
			}()
		}
	}
}

func main() {
	//s := []int{2, 3, 4, 5, 6, 7}
	//a := unsafe.Pointer(&s[0])
	//b := uintptr(a)
	//c := uintptr(a) + 4*unsafe.Sizeof(s[0])
	//d := unsafe.Pointer(c)
	//*(*int)(d) = 100
	//fmt.Println(a, b, c, d, s)
	//fmt.Println(runtime.NumCPU())
	go func() {
		var t time.Time
		for atomic.LoadInt32(&stop) == 0 {
			t = time.Now()
			runtime.GC()
			sum += time.Since(t)
			count++
		}
		fmt.Printf("GC spend avg: %v\n", time.Duration(int64(sum)/count))
	}()

	concat()
	atomic.StoreInt32(&stop, 1)
}
