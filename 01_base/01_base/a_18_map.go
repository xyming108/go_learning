package main

import "fmt"

func main() {
	var m2 map[string]int //key为string类型，value为int类型
	//初始化
	m2 = make(map[string]int, 10) //10为预估的容量
	m2["理想"] = 100
	m2["现实"] = 101
	m2["距离"] = 102
	fmt.Println(m2)
	fmt.Println(m2["现实"])
	fmt.Println()

	//约定成俗，用ok接收返回的布尔值
	v, ok := m2["杨幂"]
	if !ok {
		fmt.Println("查无此key")
	} else {
		fmt.Println(v)
	}

	//map的遍历
	for k, v := range m2 {
		fmt.Println(k, v)
	}
	for k := range m2 {
		fmt.Println(k)
	}
	for _, v := range m2 {
		fmt.Println(v)
	}

	//删除
	delete(m2, "距离")
	fmt.Println(m2)

}
