package main

import (
	"sort"
	"testing"
)

/**
Author: xym
Date: 2021/4/22 12:33
Project: Go_Learning
Description:
*/

/*压力测试*/

/*func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}

	return a / b, nil
}*/

func Benchmark_Map(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		/*Division(4, 5)*/
		a := []string{"AK5SSA1B38100666",
			"AK5SSA1B38100668",
			"AK5SSA1B38100668",
			"AK5SSA1B38100666",
			"AK5SSA1B38100668",
			"AK5SSA1B38100666",
			"AK5SSA1B38100668",
			"AK5SSA1B38100666",
			"AK5SSA1B38100666",
			"AK5SSA1B38100668",
			"AK5SSA1B38100668",
			"AK5SSA1B38100666",
			"AK5SSA1B38100668",
			"AK5SSA1B38100666",
			"AK5SSA1B38100668",
			"AK5SSA1B38100666",
			"AK5SSA1B38100666",
			"AK5SSA1B38100668",
			"AK5SSA1B38100668",
			"AK5SSA1B38100666",
			"AK5SSA1B38100668",
			"AK5SSA1B38100666",
			"AK5SSA1B38100668",
			"AK5SSA1B38100666",
			"AK5SSA1B38100666",
			"AK5SSA1B38100668",
			"AK5SSA1B38100668",
			"AK5SSA1B38100666",
			"AK5SSA1B38100668",
			"AK5SSA1B38100666",
			"AK5SSA1B38100668",
			"AK5SSA1B38100666"}

		resArr := make([]string, 0)
		tempMap := make(map[string]interface{})
		for _, val := range a {
			_, ok := tempMap[val]
			if !ok {
				resArr = append(resArr, val)
				tempMap[val] = nil
			} else {
				//fmt.Println("出现重复")
				break
			}
		}

		/*sort.Strings(a)
		length := len(a)
		for i := 1; i < length; i++ {
			if a[i-1] == a[i] {
				log.Println("sn_list里面有重复值")
				break
			}
		}*/
	}
}

func Benchmark_For(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		/*Division(4, 5)*/
		a := []string{"AK5SSA1B38100666",
			"AK5SSA1B38100668",
			"AK5SSA1B38100668",
			"AK5SSA1B38100666",
			"AK5SSA1B38100668",
			"AK5SSA1B38100666",
			"AK5SSA1B38100668",
			"AK5SSA1B38100666",
			"AK5SSA1B38100666",
			"AK5SSA1B38100668",
			"AK5SSA1B38100668",
			"AK5SSA1B38100666",
			"AK5SSA1B38100668",
			"AK5SSA1B38100666",
			"AK5SSA1B38100668",
			"AK5SSA1B38100666",
			"AK5SSA1B38100666",
			"AK5SSA1B38100668",
			"AK5SSA1B38100668",
			"AK5SSA1B38100666",
			"AK5SSA1B38100668",
			"AK5SSA1B38100666",
			"AK5SSA1B38100668",
			"AK5SSA1B38100666",
			"AK5SSA1B38100666",
			"AK5SSA1B38100668",
			"AK5SSA1B38100668",
			"AK5SSA1B38100666",
			"AK5SSA1B38100668",
			"AK5SSA1B38100666",
			"AK5SSA1B38100668",
			"AK5SSA1B38100666"}

		/*resArr := make([]string, 0)
		tempMap := make(map[string]interface{})
		for _, val := range a {
			_, ok := tempMap[val]
			if !ok {
				resArr = append(resArr, val)
				tempMap[val] = nil
			} else {
				fmt.Println("出现重复")
				break
			}
		}*/

		sort.Strings(a)
		length := len(a)
		for i := 1; i < length; i++ {
			if a[i-1] == a[i] {
				//log.Println("sn_list里面有重复值")
				break
			}
		}
	}
}

/*func Benchmark_TimeConsumingFunction(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数

	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能

	b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		Division(4, 5)
	}
}*/
