package main

import "fmt"

/*
给你一个整数数组nums ，数组中的元素互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

示例 1：
输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

示例 2：
输入：nums = [0]
输出：[[],[0]]
*/

/*
迭代法实现子集枚举
时间复杂度：O(n*(2^n))
空间复杂度：O(n)
*/
func subsets1(nums []int) [][]int {
	var ans [][]int
	n := len(nums)
	for i := 0; i < (1 << n); i++ {
		var set []int
		for j := 0; j < n; j++ {
			if ((i >> j) & 1) > 0 {
				set = append(set, nums[j])
			}
		}
		ans = append(ans, append([]int(nil), set...))
	}
	return ans
}

/*
递归法实现子集枚举
时间复杂度：O(n*(2^n))
空间复杂度：O(n)
解析：https://leetcode-cn.com/problems/subsets/solution/shou-hua-tu-jie-zi-ji-hui-su-fa-xiang-jie-wei-yun-/
*/
func subsets2(nums []int) [][]int {
	var ans [][]int
	var set []int
	n := len(nums)
	var dfs func(cur int)
	dfs = func(cur int) {
		if cur == n {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1] //回溯
		dfs(cur + 1)
	}
	dfs(0)
	return ans
}

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(subsets1(arr))
	fmt.Println(subsets2(arr))
}
