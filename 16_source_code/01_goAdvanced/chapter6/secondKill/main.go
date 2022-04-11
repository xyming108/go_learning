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

//本书重点关注后端逻辑，其余部分没有全部实现，感兴趣的读者可以自行补充实现
package main

import (
	"fmt"
	cachePkg "gitee.com/shirdonl/goAdvanced/chapter6/secondKill/cache"
	"gitee.com/shirdonl/goAdvanced/chapter6/secondKill/util"
	"github.com/garyburd/redigo/redis"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var (
	localCache  cachePkg.LocalCache
	remoteCache cachePkg.RemoteCacheKeys
	redisPool   *redis.Pool
	done        chan int
)

//初始化结构体和Redis连接池
func init() {
	localCache = cachePkg.LocalCache{
		LocalInStock:     150,
		LocalSalesVolume: 0,
	}
	remoteCache = cachePkg.RemoteCacheKeys{
		SpikeOrderHashKey:  "goods_hash_key",
		TotalInventoryKey:  "goods_total_number",
		QuantityOfOrderKey: "goods_sold_number",
	}
	redisPool = cachePkg.NewPool()
	done = make(chan int, 1)
	done <- 1
}

func main() {
	http.HandleFunc("/seckill/goods/1", seckillController)
	http.HandleFunc("/sk/view", viewHandler)
	http.HandleFunc("/sk/login", loginController)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8085", nil)
}

//处理请求函数,根据请求将响应结果信息写入日志
func seckillController(w http.ResponseWriter, r *http.Request) {
	redisConn := redisPool.Get()
	LogMsg := ""
	<-done
	//全局读写锁
	if localCache.LocalDeductionStock() && remoteCache.DeductStock(redisConn) {
		util.ResponseJson(w, 1, "抢购成功", nil)
		LogMsg = LogMsg + "result:1,localSales:" + strconv.FormatInt(localCache.LocalSalesVolume, 10)
	} else {
		util.ResponseJson(w, -1, "已售罄", nil)
		LogMsg = LogMsg + "result:0,localSales:" + strconv.FormatInt(localCache.LocalSalesVolume, 10)
	}
	//将抢购状态写入到log中
	done <- 1
	Log(LogMsg, "./record.log")
}

//处理登录逻辑
func loginController(w http.ResponseWriter, r *http.Request) {
	//...登陆逻辑
}

//秒杀页面控制器
func viewHandler(w http.ResponseWriter, r *http.Request) {
	// 1. 解析模板
	t, err := template.ParseFiles("./views/goods/secondKill.html")
	if err != nil {
		fmt.Println("template parsefile failed, err:", err)
		return
	}
	// 2.渲染模板
	name := "我爱Go语言"
	t.Execute(w, name)
}

func Log(msg string, logPath string) {
	fd, _ := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer fd.Close()
	content := strings.Join([]string{msg, "\r\n"}, "")
	buf := []byte(content)
	fd.Write(buf)
}
