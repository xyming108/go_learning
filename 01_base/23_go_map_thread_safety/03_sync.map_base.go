package main

import (
	"fmt"
	"sync"
)

func main() {
	var ma sync.Map
	//存储
	ma.Store("key1", "你好")
	if a, ok := ma.Load("key1"); ok {
		fmt.Println("====1:", a)
	}
	//删除
	ma.Delete("key1")
	if a, ok := ma.Load("key1"); ok {
		fmt.Println("====2:", a)
	}
	// LoadOrStore 获取值，如果没有则存储
	if b, ok := ma.LoadOrStore("key2", "我曾经没有"); ok {
		fmt.Println("----1:", b)
	}
	if b, ok := ma.Load("key2"); ok {
		fmt.Println("----2:", b)
	}
	//删除，有bool判断是否存在
	if b, ok := ma.LoadAndDelete("key2"); ok {
		fmt.Println("----3:", b)
	}

	ma.Store("key1", "你好!")
	ma.Store("key2", "你好!!")
	ma.Store("key3", "你好!!!")
	//无需遍历
	ma.Range(func(key, value any) bool {
		fmt.Printf("key:%s ,value:%s \n", key, value)
		//如果返回：false，则退出循环，
		return true
	})
}
