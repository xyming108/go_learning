package main

/*
给定一个 n×n 的二维矩阵matrix 表示一个图像。请你将图像顺时针旋转 90 度。
你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。

示例 1：
输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[[7,4,1],[8,5,2],[9,6,3]]

示例 2：
输入：matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
输出：[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
*/

func rotate(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return nil
	}
	xlen := len(matrix[0])
	ylen := len(matrix)

	var result [][]int
	for i := 0; i <= xlen-1; i++ {
		var tempArr []int
		for j := ylen - 1; j >= 0; j-- {
			tempArr = append(tempArr, matrix[j][i])
		}
		result = append(result, tempArr)
	}
	for i := 0; i < ylen; i++ {
		for j := 0; j < xlen; j++ {
			matrix[i][j] = result[i][j]
		}
	}
	return result
}

func main() {
	/*matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}*/
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	//matrix := [][]int{}
	rotate(matrix)
}
