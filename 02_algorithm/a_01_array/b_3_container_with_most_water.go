package main

/**
Author: xym
Date: 2021/5/23 17:44
Project: golearning
Description:
*/

/*
3. 盛最多水的容器 中等
给你 n 个非负整数 a1，a2，…，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，
垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器。

输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
示例 2：

输入：height = [1,1]
输出：1
示例 3：

输入：height = [4,3,2,1,4]
输出：16
示例 4：

输入：height = [1,2,1]
输出：2
提示：

n = height.length
2 <= n <= 3 * 104
0 <= height[i] <= 3 * 104
*/

import (
	"fmt"
)

func main() {
	fmt.Println(maxArea1([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // 49
	fmt.Println(maxArea2([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // 49
}

//
// 暴力遍历
//
func maxArea1(height []int) int {
	n := len(height)
	if n <= 1 {
		return 0
	}
	maxMulti := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			w := j - i                     // 宽
			h := min(height[i], height[j]) // 高
			maxMulti = max(maxMulti, w*h)  // 比较面积
		}
	}
	return maxMulti
}

//
// 优化1：双指针法
// 两条线段之间的面积受限于最短的线段，线段间距越长，面积越大
// 使用 2 个指针指向首部和尾部，将短指针向长指针方向移动，看能不能找到更长的线，使面积更大
// 依据：向长线方向每次移动 1 格，虽然宽度-1，但是(高度变高)*(宽度-1) >= 高度*宽度
//
func maxArea2(height []int) int {
	maxMulti := 0
	left, right := 0, len(height)-1
	for left < right {
		w := right - left
		h := min(height[left], height[right])
		maxMulti = max(maxMulti, w*h)
		if height[left] <= height[right] {
			left++ // 往右边走找更长的线
		} else {
			right-- // 往左边走
		}
	}
	return maxMulti
}
