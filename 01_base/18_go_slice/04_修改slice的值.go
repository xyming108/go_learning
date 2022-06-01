package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4}
	//这种遍历方式中，v指向的不是mySlice中的原地址，而是会重新开辟地址
	for _, v := range mySlice {
		v *= 2
	}
	fmt.Println(mySlice)

	//这种遍历方式中，用下标的方式可以修改原切片
	for i := range mySlice {
		mySlice[i] *= 2
	}
	fmt.Println(mySlice)
}
