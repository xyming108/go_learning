package main

import (
	"fmt"
)

/*
给定一个长度为 n 的整数数组height。有n条垂线，第 i 条线的两个端点是(i, 0)和(i, height[i])。
找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。
返回容器可以储存的最大水量。

说明：你不能倾斜容器。

示例 1：
输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为49。

示例 2：
输入：height = [1,1]
输出：1
*/

/*
对撞指针法
时间复杂度：O(n)
空间复杂度：O(1)
*/
func maxArea(height []int) int {
	first, second := 0, len(height)-1
	area := 0
	h := 0
	for first < second {
		tempArea := 0
		width := second - first
		if height[first] < height[second] {
			h = height[first]
			first++
		} else {
			h = height[second]
			second--
		}
		tempArea = width * h
		if tempArea > area {
			area = tempArea
		}
	}
	return area
}

func main() {
	data := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(data))
}
