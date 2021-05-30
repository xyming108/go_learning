package main

import "fmt"

//递归

//计算n的阶乘
func dg(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * dg(n-1)
}

//上台阶面试题
func taijie(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return taijie(n-1) + taijie(n-2)
}

func main() {
	ret := dg(5)
	fmt.Println(ret)

	ret1 := taijie(10)
	fmt.Println(ret1)
}
