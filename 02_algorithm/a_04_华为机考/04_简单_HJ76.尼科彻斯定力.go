package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
描述
验证尼科彻斯定理，即：任何一个整数m的立方都可以写成m个连续奇数之和

例如：
1^3=1
2^3=3+5
3^3=7+9+11
4^3=13+15+17+19

输入一个正整数m（m≤100），将m的立方写成m个连续奇数之和的形式输出

输入描述：
输入一个int整数

输出描述：
输出分解后的string

示例1
输入：
6
输出：
31+33+35+37+39+41
*/

//转化为数学问题，即等差数列求和问题，公差为2，由规律可知项数n等于输入的m
func VerifyingNicochaseTheorem(in int) string {
	sum := math.Pow(float64(in), 3)
	n := in
	a1 := int(sum)/n - (n - 1)
	var s string
	s += strconv.Itoa(a1)
	for i := 1; i < n; i++ {
		a1 = a1+ 2
		s += "+"
		s += strconv.Itoa(a1)
	}
	return s
}

func main() {
	var in int
	fmt.Scanln(&in)
	fmt.Println(VerifyingNicochaseTheorem(in))
}
