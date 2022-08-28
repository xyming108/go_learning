package main

import "fmt"

/*
给定一个仅包含0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。

示例 1：
输入：matrix = [['1','0','1','0','0'],['1','0','1','1','1'],['1','1','1','1','1'],['1','0','0','1','0']]
输出：6

示例 2：
输入：matrix = []
输出：0

示例 3：
输入：matrix = [['0']]
输出：0

示例 4：
输入：matrix = [['1']]
输出：1

示例 5：
输入：matrix = [['0','0']]
输出：0

题解：
https://leetcode.cn/problems/maximal-rectangle/solution/zui-da-ju-xing-by-leetcode-solution-bjlu/
*/

/*
单调栈
时间复杂度：O(mn)
空间复杂度：O(mn)
*/
func maximalRectangle1(matrix [][]byte) int {
	min := func(i int, i2 int) int {
		if i < i2 {
			return i
		}
		return i2
	}
	max := func(i int, i2 int) int {
		if i < i2 {
			return i2
		}
		return i
	}

	var ans int
	if len(matrix) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	left := make([][]int, m)
	for i, row := range matrix {
		left[i] = make([]int, n)
		for j, v := range row {
			if v == '0' {
				continue
			}
			if j == 0 {
				left[i][j] = 1
			} else {
				left[i][j] = left[i][j-1] + 1
			}
		}
	}
	for i, row := range matrix {
		for j, v := range row {
			if v == '0' {
				continue
			}
			width := left[i][j]
			area := width
			for k := i - 1; k >= 0; k-- {
				width = min(width, left[k][j])
				area = max(area, (i-k+1)*width)
			}
			ans = max(ans, area)
		}
	}

	return ans
}

/*
单调栈
时间复杂度：O(mn)
空间复杂度：O(mn)
*/
func maximalRectangle2(matrix [][]byte) int {
	var ans int
	max := func(i int, i2 int) int {
		if i < i2 {
			return i2
		}
		return i
	}

	if len(matrix) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	left := make([][]int, m)
	for i, row := range matrix {
		left[i] = make([]int, n)
		for j, v := range row {
			if v == '0' {
				continue
			}
			if j == 0 {
				left[i][j] = 1
			} else {
				left[i][j] = left[i][j-1] + 1
			}
		}
	}
	for j := 0; j < n; j++ { // 对于每一列，使用基于柱状图的方法
		up := make([]int, m)
		down := make([]int, m)
		stk := []int{}
		for i, l := range left {
			for len(stk) > 0 && left[stk[len(stk)-1]][j] >= l[j] {
				stk = stk[:len(stk)-1]
			}
			up[i] = -1
			if len(stk) > 0 {
				up[i] = stk[len(stk)-1]
			}
			stk = append(stk, i)
		}
		stk = nil
		for i := m - 1; i >= 0; i-- {
			for len(stk) > 0 && left[stk[len(stk)-1]][j] >= left[i][j] {
				stk = stk[:len(stk)-1]
			}
			down[i] = m
			if len(stk) > 0 {
				down[i] = stk[len(stk)-1]
			}
			stk = append(stk, i)
		}
		for i, l := range left {
			height := down[i] - up[i] - 1
			area := height * l[j]
			ans = max(ans, area)
		}
	}
	return ans
}

func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalRectangle1(matrix))
	fmt.Println(maximalRectangle2(matrix))
}
