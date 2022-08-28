package main

import "fmt"

/*
给定n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水（格子空隙处接雨水）。

示例 1：
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

示例 2：
输入：height = [4,2,0,3,2,5]
输出：9

题解：
https://leetcode.cn/problems/trapping-rain-water/solution/jie-yu-shui-by-leetcode-solution-tuvc/
*/

/*
动态规划
时间复杂度：O(n)
空间复杂度：O(n)
*/
func trap1(height []int) int {
	max := func(i int, i2 int) int {
		if i < i2 {
			return i2
		}
		return i
	}
	min := func(i int, i2 int) int {
		if i < i2 {
			return i
		}
		return i2
	}
	n := len(height)
	if n == 0 {
		return 0
	}
	leftMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}
	var ans int
	for i, v := range height {
		ans += min(leftMax[i], rightMax[i]) - v
	}
	return ans
}

/*
单调栈
时间复杂度：O(n)
空间复杂度：O(n)
*/
func trap2(height []int) int {
	min := func(i int, i2 int) int {
		if i < i2 {
			return i
		}
		return i2
	}
	var stack []int
	var ans int
	for i, v := range height {
		for len(stack) > 0 && v > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			curWidth := i - left - 1
			curHeight := min(v, height[left]) - height[top]
			ans += curWidth * curHeight
		}
		stack = append(stack, i)
	}
	return ans
}

/*
双指针
时间复杂度：O(n)
空间复杂度：O(1)
*/
func trap3(height []int) int {
	max := func(i int, i2 int) int {
		if i < i2 {
			return i2
		}
		return i
	}

	var ans int
	left, right, leftMax, rightMax := 0, len(height)-1, 0, 0
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if height[left] < height[right] {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}
	return ans
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap1(height))
	fmt.Println(trap2(height))
	fmt.Println(trap3(height))
}
