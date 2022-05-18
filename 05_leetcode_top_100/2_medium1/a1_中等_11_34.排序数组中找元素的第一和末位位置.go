package main

import (
	"fmt"
	"sort"
)

/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回[-1, -1]。

进阶：
你可以设计并实现时间复杂度为O(log n)的算法解决此问题吗？

示例 1：
输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]

示例2：
输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]

示例 3：
输入：nums = [], target = 0
输出：[-1,-1]
*/

/*
暴力遍历
时间复杂度：O(n)
空间复杂度：O(1)
*/
func searchRange1(nums []int, target int) []int {
	mapFirst := make(map[int]int)
	mapEnd := make(map[int]int)
	mapFirst[target], mapEnd[target] = -1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == target && mapFirst[target] == -1 {
			mapFirst[target] = i
		}
		if nums[i] == target {
			mapEnd[target] = i
		}
		if i < len(nums)-1 && nums[i] == target && nums[i+1] != target {
			break
		}
	}
	return []int{mapFirst[target], mapEnd[target]}
}

/*
二分查找
时间复杂度：O(log(n))
空间复杂度：O(1)
*/
func searchRange2(nums []int, target int) []int {
	//sort.SearchInts 查找第一个找到的元素的下标(从0开始计数)，如果没有找到，则返回target可以插入的位置下标
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	rightmost := sort.SearchInts(nums, target+1) - 1
	return []int{leftmost, rightmost}
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange1(nums, 8))
	fmt.Println(searchRange2(nums, 8))
}
