package main

import "fmt"

func main() {
	//copy
	a1 := []int{1, 2, 3, 4}
	a2 := a1
	var a3 = make([]int, 4, 4)
	a2[0] = 12
	fmt.Printf("a1:%x, a2:%x, a3:%x\n", &a1[0], &a2[0], &a3[0])
	fmt.Printf("a1:%x, a2:%x, a3:%x\n", &a1[1], &a2[1], &a3[1])
	fmt.Printf("a1:%x, a2:%x, a3:%x\n", &a1[2], &a2[2], &a3[2])
	fmt.Printf("a1:%x, a2:%x, a3:%x\n", &a1[3], &a2[3], &a3[3])
	fmt.Println()
	copy(a3, a1)
	fmt.Println(a1, a2, a3)
	a1[0] = 100
	fmt.Println(a1, a2, a3)

	//将a中的索引为1的这个元素删掉
	//a1 = append(a1[:1], a1[2:]...)
	//fmt.Println(a1)
	//fmt.Println(cap(a1), len(a1))
}
