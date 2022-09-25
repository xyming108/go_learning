package _4_memory_leak

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(time.Second)
	go func() {
		http.ListenAndServe("127.0.0.1:8088", nil)
	}()
	defer ticker.Stop()
	go func(ticker *time.Ticker) {
		for {
			select {
			case value := <-ticker.C:
				fmt.Println(value)
			}
		}
	}(ticker)
	time.Sleep(time.Second * 5)
	fmt.Println("finish!!!")
}
