package main

import "fmt"

/*
给你一个整数 n ，对于0 <= i <= n 中的每个 i ，计算其二进制表示中 1 的个数 ，
返回一个长度为 n + 1 的数组 ans 作为答案。

示例 1：
输入：n = 2
输出：[0,1,1]
解释：
0 --> 0
1 --> 1
2 --> 10
*/

//Brian Kernighan 算法
//时间复杂度：O(nlog(n))
//空间复杂度：O(1)
func countBits1(n int) []int {
	var onesCount func(x int) int
	onesCount = func(x int) (ones int) {
		for ; x > 0; x &= x - 1 {
			ones++
		}
		return
	}

	bits := make([]int, n+1)
	for i := range bits {
		bits[i] = onesCount(i)
	}
	return bits
}

//动态规划——最高有效位
//时间复杂度：O(n)
//空间复杂度：O(1)
func countBits2(n int) []int {
	bits := make([]int, n+1)
	highBit := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			highBit = i
		}
		bits[i] = bits[i-highBit] + 1
	}
	return bits
}

//动态规划——最低有效位
//时间复杂度：O(n)
//空间复杂度：O(1)
func countBits3(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {
		bits[i] = bits[i>>1] + i&1
	}
	return bits
}

//动态规划——最低设置位
//时间复杂度：O(n)
//空间复杂度：O(1)
func countBits4(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {
		bits[i] = bits[i&(i-1)] + 1
	}
	return bits
}

func main() {
	n := 10
	fmt.Println(countBits1(n))
	fmt.Println(countBits2(n))
	fmt.Println(countBits3(n))
	fmt.Println(countBits4(n))
}
