package main

import "fmt"

/*
给定一个二维矩阵 matrix，以下类型的多个请求：
计算其子矩形范围内元素的总和，该子矩阵的 左上角 为 (row1, col1) ，右下角 为 (row2, col2) 。
实现 NumMatrix 类：
NumMatrix(int[][] matrix) 给定整数矩阵 matrix 进行初始化
int sumRegion(int row1, int col1, int row2, int col2) 返回 左上角 (row1, col1) 、右下角 (row2, col2) 所描述的子矩阵的元素 总和 。

示例1：
输入:
["NumMatrix","sumRegion","sumRegion","sumRegion"]
[[[[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]],[2,1,4,3],[1,1,2,2],[1,2,2,4]]
输出:
[null, 8, 11, 12]
解释:
NumMatrix numMatrix = new NumMatrix([[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]);
numMatrix.sumRegion(2, 1, 4, 3); // return 8 (红色矩形框的元素总和)
numMatrix.sumRegion(1, 1, 2, 2); // return 11 (绿色矩形框的元素总和)
numMatrix.sumRegion(1, 2, 2, 4); // return 12 (蓝色矩形框的元素总和)

提示：
m == matrix.length
n == matrix[i].length
1 <= m, n <= 200
-105 <= matrix[i][j] <= 105
0 <= row1 <= row2 < m
0 <= col1 <= col2 < n
最多调用 104 次 sumRegion 方法
*/

/*
方法一：一维前缀和
先把每一个行的前缀和都记录下来，当计算二维时遍历统计每一行的子数组总和
时间复杂度：O(mn)
空间复杂度：O(mn)
*/
type NumMatrix struct {
	nums [][]int
}

/*func Constructor(matrix [][]int) NumMatrix {
	sums := make([][]int, len(matrix))
	for i, row := range matrix {
		sums[i] = make([]int, len(row)+1)
		for j, v := range row {
			sums[i][j+1] = sums[i][j] + v
		}
	}
	return NumMatrix{nums: sums}
}
func (m *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	var sum int
	for i := row1; i <= row2; i++ {
		sum += m.nums[i][col2+1] - m.nums[i][col1]
	}
	return sum
}*/

/*
方法二：二维前缀和
先把二维的前缀和都记录下来，即左上角的矩形总和，当计算二维时用右下角总和减去左上角的总和
和上面一维前缀和相比，少了一次对行的遍历
时间复杂度：O(mn)
空间复杂度：O(mn)
*/
func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	sums := make([][]int, m+1)
	sums[0] = make([]int, n+1)
	for i, row := range matrix {
		sums[i+1] = make([]int, n+1)
		for j, v := range row {
			sums[i+1][j+1] = sums[i+1][j] + v + sums[i][j+1] - sums[i][j]
		}
	}
	return NumMatrix{nums: sums}
}
func (m *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return m.nums[row2+1][col2+1] - m.nums[row2+1][col1] - m.nums[row1][col2+1] + m.nums[row1][col1]
}

func main() {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	obj := Constructor(matrix)
	param_1 := obj.SumRegion(2, 1, 4, 3)
	param_2 := obj.SumRegion(1, 1, 2, 2)
	param_3 := obj.SumRegion(1, 2, 2, 4)
	fmt.Println(param_1)
	fmt.Println(param_2)
	fmt.Println(param_3)
}
