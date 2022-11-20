package main

import "fmt"

/*
给你一个整数数组 nums 和一个整数k ，请你统计并返回 该数组中和为k的子数组的个数。

示例 1：
输入：nums = [1,1,1], k = 2
输出：2

示例 2：
输入：nums = [1,2,3], k = 3
输出：2

解法：https://leetcode.cn/problems/subarray-sum-equals-k/solution/he-wei-kde-zi-shu-zu-by-leetcode-solution/
*/

/*
暴力枚举法
时间复杂度：O(n^2)
空间复杂度：O(1)
*/
func subarraySum(nums []int, k int) int {
	var ans int
	for i := 0; i < len(nums); i++ {
		sum := 0
		j := i
		for j < len(nums) {
			sum += nums[j]
			j++
			if sum == k {
				ans++
			}
		}
	}
	return ans
}

/*
前缀和 + 哈希表优化
时间复杂度：O(n)
空间复杂度：O(n)
*/
func subarraySum2(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{0: 1}
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok {
			count += m[pre-k]
		}
		m[pre]++
	}
	return count
}

func main() {
	nums := []int{1, 1, 1}
	fmt.Println(subarraySum(nums, 2))
	fmt.Println(subarraySum2(nums, 2))
}
