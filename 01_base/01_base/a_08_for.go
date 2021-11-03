package main

import "fmt"

func main() {
	//基本格式
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i == 2 {
			break
		}
	}
	fmt.Println()

	//变形1
	var j = 5
	for ; j < 7; j++ {
		fmt.Println(j)
	}
	fmt.Println()

	//变形2
	var k = 8
	for k < 10 {
		fmt.Println(k)
		k++
	}
	fmt.Println()

	//键值循环，遍历数组、切片、字符串、map和channel
	s := "hello 世界"
	for i, v := range s {
		fmt.Printf("%d： %c\n", i, v)
	}

}
