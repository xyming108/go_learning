//++++++++++++++++++++++++++++++++++++++++
// 《Go语言高级开发与实战》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 知乎：https://www.zhihu.com/people/shirdonl
// 公众号:源码大数据
// 仓库地址：https://gitee.com/shirdonl/goAdvanced
// 仓库地址：https://github.com/shirdonl/goAdvanced
//++++++++++++++++++++++++++++++++++++++++

package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"net/http"
	"sync"
	"time"
)

type Limiter struct {
	mu        sync.Mutex    // 互斥锁（排他锁）
	limit     *rate.Limiter // 放入桶的频率  float64 类型
	burst     int           // 桶的大小
	tokens    float64       // 令牌 token 当前剩余的数量
	last      time.Time     // 最近取走 token 的时间
	lastEvent time.Time     // 最近限流事件的时间
}

type Reservation struct {
	ok        bool          // 是否满足条件分配了 token
	lim       *Limiter      // 发送令牌的限流器
	tokens    int           // 发送 token 令牌的数量
	timeToAct time.Time     // 满足令牌发放的时间
	limit     *rate.Limiter // 令牌发放速度
}

func NewLimiter(r *rate.Limiter, b int) *Limiter {
	return &Limiter{
		limit: r, // 放入桶的频率
		burst: b, // 桶的大小
	}
}

func GetApi() {
	api := "http://localhost:8081/"
	res, err := http.Get(api)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		fmt.Printf("get api success\n")
	}
}

func GetApi1() {
	api := "http://127.0.0.1:8081/"
	res, err := http.Get(api)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		fmt.Printf("get api success\n")
	}
}

//func Benchmark_Main(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		GetApi()
//	}
//}
