package _2_sync

import (
	"log"
	"sync"
	"testing"
	"time"
)

func TestRWLock(t *testing.T) {
	var wg sync.WaitGroup
	var rwLock sync.RWMutex
	flagMap := make(map[int64]int, 0)

	time1 := time.Now().Unix()
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(ii int) {
			defer wg.Done()
			for j := ii * 1000; j < (ii+1)*1000; j++ { // 给goroutine划分执行区间
				rwLock.Lock()
				flagMap[int64(j)] = 2
				//log.Println("---------该条json压缩失败或为空，已跳过, id: ", j)
				rwLock.Unlock()
			}
		}(i)
	}
	wg.Wait() // 等待所有goroutine执行完毕

	time2 := time.Now().Unix()
	log.Println("============开始时间：", time1)
	log.Println("============结束时间：", time2)
}
