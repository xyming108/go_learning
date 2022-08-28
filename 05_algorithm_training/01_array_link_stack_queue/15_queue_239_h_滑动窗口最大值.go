package main

import (
	"fmt"
)

/*
给你一个整数数组 nums，有一个大小为k的滑动窗口从数组的最左侧移动到数组的最右侧。
你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
返回 滑动窗口中的最大值
提示：1 <= k <= nums.length

示例 1：
输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]

示例 2：
输入：nums = [1], k = 1
输出：[1]
*/

/*
单调队列：（队列中的值单调递增或递减）
时间复杂度：O(n)
空间复杂度：O(k)
题解：
https://leetcode.cn/problems/sliding-window-maximum/solution/dong-hua-yan-shi-dan-diao-dui-lie-239hua-hc5u/
*/
func maxSlidingWindow(nums []int, k int) []int {
	var q []int
	push := func(i int) {
		// 如果队列不为空且当前考察元素大于等于队尾元素，则将队尾元素移除，
		// 直到队列为空或当前考察元素小于新的队尾元素
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		// 存储元素下标
		q = append(q, i)
	}

	// 先放入前k个元素下标
	for i := 0; i < k; i++ {
		push(i)
	}

	n := len(nums)
	// 窗口个数
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q[0]]
	for i := k; i < n; i++ {
		push(i)
		// 当队首元素的下标小于滑动窗口左侧边界时
		// 表示队首元素已经不再滑动窗口内，因此将其从队首移除
		for q[0] <= i-k {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
}
