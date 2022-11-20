package main

import "fmt"

/*
给你一个下标从 1 开始的整数数组 numbers ，该数组已按 非递减顺序排列，
请你从数组中找出满足相加之和等于目标数 target 的两个数。如果设这两个数分别是
numbers[index1] 和 numbers[index2] ，则 1 <= index1 < index2 <= numbers.length 。
以长度为 2 的整数数组 [index1, index2] 的形式返回这两个整数的下标 index1 和 index2。

你可以假设每个输入 只对应唯一的答案 ，而且你 不可以 重复使用相同的元素。
你所设计的解决方案必须只使用常量级的额外空间。
和上一个两数之和题目是一样的，只是这个输入的数组是有顺序的，既然有顺序，那么可以得到时间空间复杂度更低的解法

示例 1：
输入：numbers = [2,7,11,15], target = 9
输出：[1,2]
解释：2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。返回 [1, 2] 。

示例 2：
输入：numbers = [2,3,4], target = 6
输出：[1,3]
解释：2 与 4 之和等于目标数 6 。因此 index1 = 1, index2 = 3 。返回 [1, 3] 。

示例 3：
输入：numbers = [-1,0], target = -1
输出：[1,2]
解释：-1 与 0 之和等于目标数 -1 。因此 index1 = 1, index2 = 2 。返回 [1, 2] 。
*/

/*
二分查找
每次遍历数组固定左边的下标，然后在右边的下标中采用二分查找的方式寻找第二个下标
时间复杂度：O(nlog(n))
空间复杂度：O(1)
*/
func twoSum3(numbers []int, target int) []int {
	for i, v := range numbers {
		l, r := i+1, len(numbers)-1
		for l <= r {
			m := (l + r) / 2
			if v+numbers[m] == target {
				return []int{i + 1, m + 1}
			} else if v+numbers[m] < target {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}
	return []int{-1, -1}
}

/*
双指针
时间复杂度：O(n)
空间复杂度：O(1)
*/
func twoSum4(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		} else if sum < target {
			l++
		} else {
			r--
		}
	}
	return []int{-1, -1}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result3 := twoSum3(nums, target)
	fmt.Println(result3)
	nums = []int{2, 7, 11, 15}
	result4 := twoSum4(nums, target)
	fmt.Println(result4)
}
