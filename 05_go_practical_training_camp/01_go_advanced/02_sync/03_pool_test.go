package _2_sync

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	// 只要没有put值进来，就会new一个默认值，put进来的值只要被get了一次就不在pool里面了
	pool := sync.Pool{
		New: func() interface{} {
			fmt.Println("新建pool")
			return "我是默认值"
		},
	}

	a1 := pool.Get().(string)
	fmt.Println("=====a1:", a1)

	a2 := pool.Get().(string)
	fmt.Println("=====a2:", a2)

	fmt.Println("--------------------")

	pool.Put("我放了一个值进来")

	a3 := pool.Get().(string)
	fmt.Println("=====a3:", a3)

	a4 := pool.Get().(string)
	fmt.Println("=====a4:", a4)

	pool.Put("上一个值被取走了，我又放了一个值进来1")
	pool.Put("上一个值被取走了，我又放了一个值进来2")

	// 遵循先进先出的原则
	a5 := pool.Get().(string)
	fmt.Println("=====a5:", a5)
	a6 := pool.Get().(string)
	fmt.Println("=====a6:", a6)
}

func TestOnce2(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "默认值"
		},
	}

	//了解Get和Put的特性
	fmt.Println(pool.Get())	// 没有放入自定义的值就会返回默认值
	pool.Put("新值")      // 放入自定义的值
	fmt.Println(pool.Get())	// 取出刚才放入的值
	fmt.Println(pool.Get()) // 放入的值取完了，返回默认值
	fmt.Println()

	//基本使用
	pool.Put("新值2")		// 放入想要复用的值
	pool.Put("新值3")		// 放入想要复用的值
	pool.Put("新值4")		// 放入想要复用的值
	fmt.Println(pool.Get())	// 取出使用
	fmt.Println(pool.Get())	// 取出使用
	fmt.Println(pool.Get())	// 取出使用
	pool.Put("新值2") // 放回去给下次使用
}
