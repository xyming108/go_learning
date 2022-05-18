package main

import (
	"fmt"
)

/*
给定一个非负整数数组nums ，你最初位于数组的 第一个下标 。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个下标。

示例1：
输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。

示例2：
输入：nums = [3,2,1,0,4]
输出：false
解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。
*/

/*
贪心
时间复杂度：O(n)
空间复杂度：O(1)
*/
func canJump(nums []int) bool {
	var maxNum func(i, j int) int
	maxNum = func(i, j int) int {
		if i > j {
			return i
		} else {
			return j
		}
	}

	n := len(nums)
	rightmost := 0
	for i := 0; i < n; i++ {
		if i <= rightmost {
			rightmost = maxNum(i+nums[i], rightmost)
			if rightmost >= n-1 {
				return true
			}
		}
	}

	return false
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(nums))
}
