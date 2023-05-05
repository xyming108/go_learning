package main

import (
	"go_Learning/17_word_compare/speedControl"
	"go_Learning/17_word_compare/util"
	"log"
	"strconv"
	"time"
)

func init() {
	util.InitChannel()
	util.InitLarkChannel()
}

func main() {
	//t1 := time.Now().Format("2006-01-02 15:04:05")
	//wordCompare.WordCompare1()
	//wordCompare.WordCompare2()
	//t2 := time.Now().Format("2006-01-02 15:04:05")
	//fmt.Println("开始时间: ", t1)
	//fmt.Println("结束时间: ", t2)

	// 模拟前端请求
	tick := time.Tick(time.Millisecond*500)
	var i int
	go func() {
		for {
			select {
			case <-tick:
				go speedControl.SendMsgToChan("sn" + strconv.Itoa(i))
				i++
			}
		}
	}()

	//speedControl.SendMsgToChan("sn")

	// 先把所有sku+position组合置为false
	keyMap := make(map[string]int, 0)
	for i := 0; i < 10; i++ {
		key := strconv.Itoa(i) + "ABC" + strconv.Itoa(i)
		speedControl.SafeMap.Store(key, false)
		keyMap[key] = 1
	}

	// 消费
	go func() {
		err := speedControl.SendToLark()
		if err != nil {
			log.Fatal("------------2:", err)
			return
		}
	}()

	// 发送
	err := speedControl.SendToLarkChan()
	if err != nil {
		log.Fatal("------------1:", err)
		return
	}

	// 判断
	//length := len(keyMap)
	//for length > 0 {
	//	for k := range keyMap {
	//		if flag, ok := speedControl.SafeMap.Load(k); ok {
	//			if flag == true {
	//				length--
	//				fmt.Println("接收************************length: ", length, " ", k)
	//				speedControl.SafeMap.Delete(k)
	//			}
	//		}
	//	}
	//}
	//
	//if length == 0 {
	//	fmt.Println("++++++++++++++++++++++: 上面图片数据落地数据库完成，开始从数据库拿数据并执行文字识别")
	//	speedControl.SafeMap.Range(func(key, value any) bool {
	//		fmt.Println("  key: ", key)
	//		fmt.Println("value: ", value)
	//		return true
	//	})
	//}
}
