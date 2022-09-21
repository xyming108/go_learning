package _2_sync

import (
	"fmt"
	"sync"
	"testing"
)

func TestOnce(t *testing.T) {
	var one sync.Once
	for i := 0; i < 10; i++ {
		one.Do(func() {
			fmt.Println("我只执行一次")
		})
	}
}
