package main

import "fmt"

//拥有相同类型元素的可变长度序列
//切片是引用类型，只要底层数组值改了，则所切到的数值也会改变

func main() {
	//切片的定义
	var s1 []int
	var s2 []string
	fmt.Println(s1, s2)

	//nil表示空
	fmt.Println(s1 == nil) //true

	//初始化
	s1 = []int{1, 2}
	s2 = []string{"江西", "广东", "上海"}
	fmt.Println(s1, s2)
	fmt.Println()

	//cap表示容量，指底层数组从第一个存在的元素开始计算到最后的容量
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))

	//由数组得到切片
	a1 := [...]int{1, 2, 3, 5, 6, 99}
	a2 := a1[0:3] //左闭右开
	fmt.Println(a2)

}
