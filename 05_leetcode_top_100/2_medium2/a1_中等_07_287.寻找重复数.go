package main

import "fmt"

/*
给定一个包含n + 1 个整数的数组nums ，其数字都在[1, n]范围内（包括 1 和 n），可知至少存在一个重复的整数。
假设 nums 只有 一个重复的整数 ，返回这个重复的数 。
你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。

示例 1：
输入：nums = [1,3,4,2,2]
输出：2

示例 2：
输入：nums = [3,1,3,4,2]
输出：3

示例 3：
输入：nums = [2,2,2,2,2]
输出：2

提示：
nums 中 只有一个整数 出现 两次或多次 ，其余整数均只出现 一次
*/

//题解：https://leetcode.cn/problems/find-the-duplicate-number/solution/xun-zhao-zhong-fu-shu-by-leetcode-solution/
//如果不限制空间复杂度，可以用hash的方式查重

/*
快慢指针法
时间复杂度：O(n)
空间复杂度：O(1)
慢指针走一步，快指针走两步；先让快慢指针一起走，最终会相遇，此时快指针和慢指针都在圈里的某个相同位置，
然后重新把慢指针放在开始位置，此时快慢指针每次都走一步，最终相遇的位置就是重复的数，即环的入口位置
*/
func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	slow, fast = nums[slow], nums[nums[fast]]
	for slow != fast {
		slow, fast = nums[slow], nums[nums[fast]]
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

func main() {
	nums := []int{3, 1, 3, 4, 2}
	fmt.Println(findDuplicate(nums))
}
