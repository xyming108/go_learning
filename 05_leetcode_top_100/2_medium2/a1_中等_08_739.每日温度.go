package main

import (
	"fmt"
	"math"
)

/*
给定一个整数数组temperatures，表示每天的温度，返回一个数组answer，其中answer[i]是指对于第 i 天，下一个更高温度出现在几天后。
如果气温在这之后都不会升高，请在该位置用0 来代替。

30 <= temperatures[i] <= 100

示例 1:
输入: temperatures = [73,74,75,71,69,72,76,73]
输出:[1,1,4,2,1,1,0,0]

示例 2:
输入: temperatures = [30,40,50,60]
输出:[1,1,1,0]

示例 3:
输入: temperatures = [30,60,90]
输出: [1,1,0]
*/

/*
暴力
时间复杂度：O(nm)，n是温度列表的长度，m是next的长度，
空间复杂度：O(m)
从后往前遍历，在后面这些走过的数据当中挑一个符合要求且距离最近的数，下标相减就是要求的距离
*/
func dailyTemperatures1(temperatures []int) []int {
	length := len(temperatures)
	ans := make([]int, length)
	next := make([]int, 101)
	for i := 0; i < 101; i++ {
		next[i] = math.MaxInt32
	}
	for i := length - 1; i >= 0; i-- {
		warmerIndex := math.MaxInt32
		for t := temperatures[i] + 1; t <= 100; t++ {
			if next[t] < warmerIndex {
				warmerIndex = next[t]
			}
		}
		if warmerIndex < math.MaxInt32 {
			ans[i] = warmerIndex - i
		}
		next[temperatures[i]] = i
	}
	return ans
}

/*
单调栈
时间复杂度：O(n)
空间复杂度：O(n)
*/
func dailyTemperatures2(temperatures []int) []int {
	length := len(temperatures)
	ans := make([]int, length)
	var stack []int
	for i := 0; i < length; i++ {
		temperature := temperatures[i]
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}
	return ans
}

func main() {
	temperatrues := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures1(temperatrues))
	fmt.Println(dailyTemperatures2(temperatrues))
}
