package _4_memory_leak

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestSend1(t *testing.T) {
	ch := make(chan int)
	fmt.Println("num of go start: ", runtime.NumGoroutine())
	time.Sleep(time.Second)

	for i := 0; i < 5; i++ {
		go func(ii int) {
			ch <- ii
			fmt.Println("send to chan: ", ii)
		}(i)
	}

	go func() {
		value := <-ch
		fmt.Println("recv from chan: ", value)
	}()

	time.Sleep(time.Second)
	fmt.Println("num of go end: ", runtime.NumGoroutine())
}

func TestSend2(t *testing.T) {
	ch := make(chan int, 2)
	fmt.Println("num of go start: ", runtime.NumGoroutine())
	time.Sleep(time.Second)

	for i := 0; i < 5; i++ {
		go func(ii int) {
			ch <- ii
			fmt.Println("send to chan: ", ii)
		}(i)
	}

	go func() {
		value := <-ch
		fmt.Println("recv from chan: ", value)
	}()

	time.Sleep(time.Second)
	fmt.Println("num of go end: ", runtime.NumGoroutine())
}

func TestRecv1(t *testing.T) {
	ch := make(chan int)
	fmt.Println("num of go start: ", runtime.NumGoroutine())
	time.Sleep(time.Second)

	go func() {
		ch <- 1
		fmt.Println("send to chan")
	}()

	for i := 0; i < 5; i++ {
		go func() {
			value := <-ch
			fmt.Println("recv from chan: ", value)
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("num of go end: ", runtime.NumGoroutine())
}

func TestRecv2(t *testing.T) {
	ch := make(chan int)
	fmt.Println("num of go start: ", runtime.NumGoroutine())
	time.Sleep(time.Second)

	for i := 0; i < 3; i++ {
		go func() {
			ch <- 1
			fmt.Println("send to chan")
		}()
	}

	for i := 0; i < 5; i++ {
		go func() {
			value := <-ch
			fmt.Println("recv from chan: ", value)
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("num of go end: ", runtime.NumGoroutine())
}

func TestNil(t *testing.T) {
	var ch chan int
	fmt.Println("num of go start: ", runtime.NumGoroutine())
	time.Sleep(time.Second)

	go func() {
		ch <- 1
		fmt.Println("send to chan")
	}()

	go func() {
		value := <-ch
		fmt.Println("recv from chan", value)
	}()

	time.Sleep(time.Second)
	fmt.Println("num of go end: ", runtime.NumGoroutine())
}
