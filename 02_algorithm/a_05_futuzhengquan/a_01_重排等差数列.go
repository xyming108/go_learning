package main

import (
	"fmt"
	"math"
)

/*
任意相邻两项的差都相等的数列叫做等差数列。给你一个不包含重复元素的 int 类型的数组 a ，请写一个函数，
如果数组中所有的元素能重新排列为一个等差数列，返回 true ，否则返回 false 。

例如数组［3,5,1］能重排成一个等差数列1,3,5}；而数组［0,1,3］不能重排成等差数列。

请尽可能使算法的时间和空间复杂度达到最优。
*/

func sortArray(s []int) bool {
	min, max := math.MaxInt, math.MinInt
	for i := 0; i < len(s); i++ {
		if s[i] > max {
			max = s[i]
		}
		if s[i] < min {
			min = s[i]
		}
	}

	tempMap := make(map[int]int, len(s))
	for i := 0; i < len(s); i++ {
		if _, ok := tempMap[s[i]]; !ok {
			tempMap[s[i]] = s[i]
		}
	}

	var tempArr []int
	for j := min; j <= max; j++ {
		if v, ok := tempMap[j]; ok {
			tempArr = append(tempArr, v)
		}
	}

	for i := 1; i < len(tempArr)-1; i++ {
		if tempArr[i+1]-tempArr[i] != tempArr[i]-tempArr[i-1] {
			return false
		}
	}
	return true
}

func main() {
	s := []int{3, 5, 1}
	fmt.Println(sortArray(s))
}
