package main

import "fmt"

/*
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1
求在该柱状图中，能够勾勒出来的矩形的最大面积。

示例 1:
输入：heights = [2,1,5,6,2,3]
输出：10
解释：最大的矩形为图中红色区域，面积为 10

示例 2：
输入：heights = [2,4]
输出：4
*/

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

/*
方法一：暴力法(会超出时间限制)
时间复杂度：O(N^2)
空间复杂度：O(1)

遍历每个矩形，依次向两边扩展，找到大于当前矩形高度的最左边和最右边的矩形下标，
然后计算相乘，求出最大高度（用时间换空间）
*/
func largestRectangleArea(heights []int) int {
	length := len(heights)
	if length == 0 {
		return 0
	}
	result := 0
	for i := 0; i < length; i++ {
		// 往左边找
		left, curHeight := i, heights[i]
		for left > 0 && heights[left-1] >= curHeight {
			left--
		}

		// 往右边找
		right := i
		for right < length-1 && heights[right+1] >= curHeight {
			right++
		}
		width := right - left + 1
		result = max(result, width*curHeight)
	}
	return result
}

/*
解法二：单调栈
时间复杂度：O(N)
空间复杂度：O(N)

详解参考：
https://leetcode.cn/problems/largest-rectangle-in-histogram/solution/bao-li-jie-fa-zhan-by-liweiwei1419/
*/
func largestRectangleArea2(heights []int) int {
	length := len(heights)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return heights[0]
	}
	result := 0
	newHeights := make([]int, len(heights)+2, len(heights)+2)
	for i, v := range heights {
		newHeights[i+1] = v
	}
	newHeights[0], newHeights[length+1] = 0, 0
	length += 2

	stack := make([]int, length, length)
	stack = append(stack, 0)
	for i := 1; i < length; i++ {
		for newHeights[i] < newHeights[stack[len(stack)-1]] {
			curHeight := newHeights[stack[len(stack)-1]]
			stack = stack[0 : len(stack)-1]
			curWidth := i - stack[len(stack)-1] - 1
			result = max(result, curWidth*curHeight)
		}
		stack = append(stack, i)
	}

	return result
}

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(heights))
	fmt.Println(largestRectangleArea2(heights))
}
