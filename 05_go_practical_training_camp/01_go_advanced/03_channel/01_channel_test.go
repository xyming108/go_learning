package _3_channel

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestChannelNoBuffer(t *testing.T) {
	ch1 := make(chan int)
	go func() {
		time.Sleep(time.Second)
		val1 := <-ch1
		fmt.Println(val1)
	}()
	go func() {
		time.Sleep(time.Second)
		val1 := <-ch1
		fmt.Println(val1)
	}()
	go func() {
		time.Sleep(time.Second)
		val1 := <-ch1
		fmt.Println(val1)
	}()
	// 如果没有用goroutine去接收通道内的值，这一步将会阻塞
	// 所以goroutine需要写在阻塞这一步的前面
	ch1 <- 10
	//close(ch1)
}

func TestWithBuffer(t *testing.T) {
	ch2 := make(chan string, 2)
	for i := 0; i < 2; i++ {
		ch2 <- "value " + strconv.Itoa(i)
	}
	for i := 0; i < 2; i++ {
		temp := <-ch2
		fmt.Println(temp)
	}
}

func TestCLose(t *testing.T) {
	cl := make(chan bool)

	//var wg sync.WaitGroup
	//wg.Add(1)
	go func() {
		//defer wg.Done()
		time.Sleep(time.Second)
		close(cl)
		rec := <-cl
		fmt.Println(rec)
	}()
	//wg.Wait()

	cl <- true
}

