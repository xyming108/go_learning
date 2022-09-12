package _2_sync

import (
	"fmt"
	"sync"
	"testing"
)

func TestDeadLock(t *testing.T) {
	var a sync.Mutex

	a.Lock()
	a.Lock()
	fmt.Println("++++++:", 1)
	a.Unlock()
	a.Unlock()
}

/*type SafeMap[K comparable, V any] struct {
	m     map[K]V
	mutex sync.RWMutex
}

//业务场景，若已经有key则直接返回，若没有则把value放进去
func (s *SafeMap[K, V]) doubleCheck(key K, value V) V {
	s.mutex.RLock()
	oldV, ok := s.m[key]
	s.mutex.RUnlock()
	if ok { //第一次检查，因为是读操作，若加读写锁能读出来，就没必要加写锁，降低开销
		return oldV
	}

	//若上面没读取出来，才会执行这步
	s.mutex.Lock()
	defer s.mutex.Unlock()
	oldV, ok = s.m[key]
	if ok { //第二次检查，也叫double-check，若第二次查出来了直接返回
		return oldV
	}
	s.m[key] = value //若第二次还没读到，就赋新值给这个key

	return value
}*/
