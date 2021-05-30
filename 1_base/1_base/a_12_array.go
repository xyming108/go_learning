package main

import (
	"fmt"
)

//数组，存放元素的容器，长度不可变
//必须指定存放的元素的类型和容量
//数组的长度是数组类型的一部分

func main() {
	var a1 [3]bool
	var a2 [4]bool
	//初始化
	a1 = [3]bool{true, false, true}
	fmt.Printf("%T %T\n", a1, a2)
	fmt.Println(a1)
	//初始化
	a3 := [3]int{2, 4, 1}
	//根据初始值自动推断数组长度的大小
	a4 := [...]int{1, 5, 7, 3}
	//指定索引
	a5 := [5]int{0: 1, 4: 8}
	fmt.Println(a3, a4, a5)

	//数组的遍历
	city := [...]string{"北京", "上海", "深圳"}
	for i := 0; i < len(city); i++ {
		fmt.Println(city[i])
	}
	for i, v := range city {
		fmt.Println(i, v)
	}

	//多维数组
	var a11 [3][2]int
	a11 = [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(a11)
	//多维数组的遍历
	for _, v1 := range a11 {
		//fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	//数组是值传递
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1, b2)

}
