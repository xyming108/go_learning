package main

import "fmt"

/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。

示例 1：
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
*/

//暴力解法
//时间复杂度：O(N²)
//空间复杂度：O(1)
func twoSum1(nums []int, target int) []int {
	for i, v := range nums{
		for j := i + 1; j < len(nums); j++ {
			sum := v + nums[j]
			if sum == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

//查找表法
//用map记录已经遍历过的数值和下标
//空间换时间
//时间复杂度：O(N)
//空间复杂度：O(N)
func twoSum2(nums []int, target int) []int {
	//mapTemp := map[int]int{}
	mapTemp := make(map[int]int, len(nums))	//初始化一定内存，内存消耗更少
	for i, v := range nums {
		if j, ok := mapTemp[target-v]; ok {
			return []int{j, i}
		}
		mapTemp[v] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result1 := twoSum1(nums, target)
	fmt.Println(result1)

	result2 := twoSum2(nums, target)
	fmt.Println(result2)
}
