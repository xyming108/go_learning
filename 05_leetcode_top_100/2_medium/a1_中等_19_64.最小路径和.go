package main

import (
	"fmt"
	"math"
)

/*
给定一个包含非负整数的 mxn网格grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
说明：每次只能向下或者向右移动一步。
注意：1 <= m, n <= 200

示例 1：
输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7
解释：因为路径 1→3→1→1→1 的总和最小。

示例 2：
输入：grid = [[1,2,3],[4,5,6]]
输出：12
*/

/*
动态规划
转移方程：
当 i>0 且 j=0 时，dp[i][0]=dp[i-1][0]+grid[i][0]
当 i=0 且 j>0 时，dp[0][j]=dp[0][j-1]+grid[0][j]
当 i>0 且 j>0 时，dp[i][j]=min(dp[i-1][j],dp[i][j-1])+grid[i][j]

时间复杂度：O(mn)
空间复杂度：O(mn)
*/

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	raw, column := len(grid), len(grid[0])
	dp := make([][]int, raw)
	for i := 0; i < raw; i++ {
		dp[i] = make([]int, column)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < raw; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < column; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for i := 1; i < raw; i++ {
		for j := 1; j < column; j++ {
			dp[i][j] = int(math.Min(float64(dp[i-1][j]), float64(dp[i][j-1]))) + grid[i][j]
		}
	}
	return dp[raw-1][column-1]
}

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}
