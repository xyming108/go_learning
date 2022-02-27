package main

import (
	"fmt"
)

/*
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例 1：
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6

示例 2：
输入：nums = [1]
输出：1
*/

//动态规划
//f(i)表示以第i个数结尾的最大连续和
//f(i)=max{f(i−1)+nums[i],nums[i]}
//时间复杂度：O(n)，其中 n 为 nums 数组的长度，我们只需要遍历一遍数组即可求得答案
//空间复杂度：O(1)，我们只需要常数空间存放若干变量
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		/*
		凡是会让总数变小的值，即<0的值，一律丢弃，
		这里也可以写成：
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		*/
		if nums[i-1]+nums[i] > nums[i] {
			nums[i] += nums[i-1]
		}
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	result := maxSubArray(nums)
	fmt.Println(result)
}
