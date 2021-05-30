package main

import "fmt"

/**
Author: xym
Date: 2021/5/23 15:04
Project: GO_Learning
Description:
*/

/*
1. 两数之和 简单
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。

示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：

输入：nums = [3,3], target = 6
输出：[0,1]
提示：

2 <= nums.length <= 103
-109 <= nums[i] <= 109
-109 <= target <= 109
只会存在一个有效答案

*/

func main() {
	fmt.Println(twoSum1([]int{2, 7, 11, 15}, 9)) // [0,1]
	fmt.Println(twoSum2([]int{3, 2, 4}, 6))      // [1,2]
	fmt.Println(twoSum3([]int{3, 3}, 6))         // [0,1]
}

//
// 暴力双重遍历
//
func twoSum1(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ { // 注意 j 的遍历区间
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

//
// 哈希表索引
// 检查数组中是否存在目标元素，若存在则找出索引
// 哈希表特别适合抽象类的配对结构，当要解决问题的数据单元是成对数据关系时，考虑哈希表 map 结构
//
func twoSum2(nums []int, target int) []int {
	num2index := make(map[int]int, len(nums))
	for i, num := range nums {
		num2index[num] = i
	}
	for i, num := range nums {
		pair := target - num
		if j, ok := num2index[pair]; ok && i != j { // 剔除自身相加的情况，使用哈希表需要注意的点：索引重叠时为同一元素
			return []int{i, j}
		}
	}
	return nil
}

//
// 优化后的哈希表索引方式
// 看 twoSum2 会发现进行了2次for循环，可以进行合并优化，一边遍历，一边检查
//
func twoSum3(nums []int, target int) []int {
	num2index := make(map[int]int, len(nums))
	for i, num := range nums {
		pair := target - num
		if j, ok := num2index[pair]; ok && i != j {
			return []int{j, i} // 注意返回值顺序，向后遍历 nums，所以 i 在 j 后
		}
		num2index[num] = i
	}
	return nil
}
