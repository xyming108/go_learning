package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"runtime/trace"
	"sync"
)

var (
	wg sync.WaitGroup
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		return
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		return
	}

	defer trace.Stop()

	runtime.GOMAXPROCS(runtime.NumCPU())
	debug.SetMaxThreads(10)
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("你好")
	}()
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()

	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.NumCgoCall())
	fmt.Println(runtime.NumCPU())
}
