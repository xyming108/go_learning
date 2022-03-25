package main

import (
	"fmt"
	"math/bits"
)

/*
两个整数之间的 汉明距离 指的是这两个数字对应二进制位不同的位置的数目。
给你两个整数 x 和 y，计算并返回它们之间的汉明距离。

示例 1：
输入：x = 1, y = 4
输出：2
解释：
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑
上面的箭头指出了对应二进制位不同的位置。

示例 2：
输入：x = 3, y = 1
输出：1
*/

//内置位计数功能
//时间复杂度：O(1)
//空间复杂度：O(1)
func hammingDistance1(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}

//移位实现位计数
//时间复杂度：O(logC)
//空间复杂度：O(1)
func hammingDistance2(x int, y int) int {
	var ans int
	for s := x ^ y; s > 0; s >>= 1 {
		ans += s & 1
	}
	return ans
}

//Brian Kernighan 算法
//时间复杂度：O(logC)
//空间复杂度：O(1)
func hammingDistance3(x int, y int) int {
	var ans int
	for s := x ^ y; s > 0; s &= s - 1 {
		ans++
	}
	return ans
}

func main() {
	x := 1
	y := 4
	fmt.Println(hammingDistance1(x, y))
	fmt.Println(hammingDistance2(x, y))
	fmt.Println(hammingDistance3(x, y))
}
