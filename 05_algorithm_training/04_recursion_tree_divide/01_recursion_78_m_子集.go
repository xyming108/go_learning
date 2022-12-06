package main

import "fmt"

/*
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

示例 1：
输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

示例 2：
输入：nums = [0]
输出：[[],[0]]

提示：
1 <= nums.length <= 10
-10 <= nums[i] <= 10
nums 中的所有元素 互不相同
*/

/*
递归，回溯法
时间复杂度：O(n*(2^n))
空间复杂度：O(n)
*/
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	var recursion func(index int, num []int)
	recursion = func(index int, num []int) {
		if index == len(nums) {
			temp := make([]int, len(num))
			copy(temp, num)
			res = append(res, temp)
			return
		}
		num = append(num, nums[index])
		recursion(index+1, num)   //选择该值往下走
		num = num[:len(num)-1]
		recursion(index+1, num)   //不选择该值往下走，回溯
	}
	recursion(0, []int{})
	return res
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}
