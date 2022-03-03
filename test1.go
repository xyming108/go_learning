package main

import (
	"fmt"
	"runtime"
)

func main() {
	done := false

	go func() {
		done = true
	}()

	for !done {
		fmt.Println("1")
		runtime.Gosched()
		fmt.Println("2")
	}

	println("done !")
}
