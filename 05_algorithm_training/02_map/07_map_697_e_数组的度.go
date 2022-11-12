package main

import (
	"fmt"
)

/*
给定一个非空且只包含非负数的整数数组 nums，数组的 度 的定义是指数组里任一元素出现频数的最大值。
你的任务是在 nums 中找到与 nums 拥有相同大小的度的最短连续子数组，返回其长度。

示例 1：
输入：nums = [1,2,2,3,1]
输出：2
解释：
输入数组的度是 2 ，因为元素 1 和 2 的出现频数最大，均为 2 。
连续子数组里面拥有相同度的有如下所示：
[1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
最短连续子数组 [2, 2] 的长度为 2 ，所以返回 2 。

示例 2：
输入：nums = [1,2,2,3,1,4,2]
输出：6
解释：
数组的度是 3 ，因为元素 2 重复出现 3 次。
所以 [2,2,3,1,4,2] 是最短子数组，因此返回 6 。

提示：
nums.length 在 1 到 50,000 范围内。
nums[i] 是一个在 0 到 49,999 范围内的整数。
*/

/*
我的解法map：
时间复杂度较高
空间复杂度：O(n)
*/
func findShortestSubArray(nums []int) int {
	numMap, maxNum, maxCount := make(map[int]int, 0), 0, 0
	for _, v := range nums {
		numMap[v]++
		if numMap[v] > maxCount {
			maxNum = v
			maxCount = numMap[v]
		}
	}

	maxSlice := make([]int, 0)
	maxSlice = append(maxSlice, maxNum)
	for k, v := range numMap {
		if v == maxCount && k != maxNum {
			maxSlice = append(maxSlice, k)
		}
	}
	minLen := 50000
	for _, v := range maxSlice {
		i, j := 0, len(nums)-1
		for i < len(nums) {
			if nums[i] == v {
				break
			}
			i++
		}
		for j >= 0 && j >= i {
			if nums[j] == v {
				break
			}
			j--
		}
		if j-i+1 < minLen {
			minLen = j - i + 1
		}
	}
	return minLen
}

/*
官方思路map
时间复杂度：O(n)
空间复杂度：O(n)
*/
type assist struct {
	count int
	l, r  int
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func findShortestSubArray2(nums []int) int {
	numMap := make(map[int]*assist, 0)
	for i, v := range nums {
		if value, ok := numMap[v]; ok {
			value.count++
			value.r = i
			numMap[v] = value
		} else {
			numMap[v] = &assist{
				count: 1,
				l:     i,
				r:     i,
			}
		}
	}

	maxCount, res := 0, 0
	for _, v := range numMap {
		if v.count > maxCount {
			maxCount = v.count
			res = v.r - v.l + 1
		} else if v.count == maxCount {
			res = min(res, v.r-v.l+1)
		}
	}
	return res
}

func main() {
	nums := []int{1, 2, 2, 3, 1, 4, 2}
	//nums := []int{1, 2, 2, 3, 1}
	fmt.Println(findShortestSubArray(nums))
	fmt.Println(findShortestSubArray2(nums))
}
