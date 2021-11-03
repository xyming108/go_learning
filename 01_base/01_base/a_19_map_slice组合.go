package main

import "fmt"

//map和slice组合

func main() {
	//元素类型为map的切片
	var s1 = make([]map[int]string, 5, 10)
	//对map初始化
	s1[0] = make(map[int]string, 1)
	s1[0][1] = "A"
	s1[0][2] = "B"
	s1[1] = make(map[int]string, 1)
	s1[1][0] = "C"
	fmt.Println(s1)

	//值为切片类型的map
	var m1 = make(map[string][]int, 10)
	//初始化
	m1["北京"] = []int{10, 20, 30}
	fmt.Println(m1)
}
