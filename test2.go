package main

import "sync"

// 代表互斥锁
type FooMap struct {
	sync.Mutex
	data map[int]int
}

// 代表读写锁
type BarRwMap struct {
	sync.RWMutex
	data map[int]int
}

var fooMap *FooMap
var barRwMap *BarRwMap
var syncMap *sync.Map

// 初始化基本数据结构
func init() {
	fooMap = &FooMap{data: make(map[int]int, 100)}
	barRwMap = &BarRwMap{data: make(map[int]int, 100)}
	syncMap = &sync.Map{}
}

func main() {

}
