package main

import "fmt"

/*
给你一个由'1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。

示例 1：
输入：grid = [
  ['1','1','1','1','0'],
  ['1','1','0','1','0'],
  ['1','1','0','0','0'],
  ['0','0','0','0','0']
]
输出：1

示例 2：
输入：grid = [
  ['1','1','0','0','0'],
  ['1','1','0','0','0'],
  ['0','0','1','0','0'],
  ['0','0','0','1','1']
]
输出：3
*/

/*
深度优先搜索
时间复杂度：O(m*n)
时间复杂度：O(m*n)
*/
func numIslands1(grid [][]byte) int {
	length := len(grid)
	if length == 0 {
		return 0
	}
	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				result++
				dfs(grid, i, j)
			}
		}
	}

	return result
}

func dfs(grid [][]byte, i, j int) {
	row, column := len(grid), len(grid[0])

	grid[i][j] = '0'
	if i-1 >= 0 && grid[i-1][j] == '1' {
		dfs(grid, i-1, j)
	}
	if i+1 < row && grid[i+1][j] == '1' {
		dfs(grid, i+1, j)
	}
	if j-1 >= 0 && grid[i][j-1] == '1' {
		dfs(grid, i, j-1)
	}
	if j+1 < column && grid[i][j+1] == '1' {
		dfs(grid, i, j+1)
	}
}

/*
广度优先搜索
时间复杂度：O(m*n)
时间复杂度：O(min(m,n))
*/
func numIslands2(grid [][]byte) int {
	length := len(grid)
	if length == 0 {
		return 0
	}
	row, column := len(grid), len(grid[0])
	result := 0
	type pair struct{ i, j int }
	for r := 0; r < row; r++ {
		for c := 0; c < column; c++ {
			if grid[r][c] == '1' {
				result++
				grid[r][c] = '0'
				var queue []pair
				queue = append(queue, pair{i: r, j: c})
				for len(queue) > 0 {
					rc := queue[len(queue)-1:]
					queue = queue[0 : len(queue)-1]
					i, j := rc[0].i, rc[0].j
					if i-1 >= 0 && grid[i-1][j] == '1' {
						queue = append(queue, pair{i: i - 1, j: j})
						grid[i-1][j] = '0'
					}
					if i+1 < row && grid[i+1][j] == '1' {
						queue = append(queue, pair{i: i + 1, j: j})
						grid[i+1][j] = '0'
					}
					if j-1 >= 0 && grid[i][j-1] == '1' {
						queue = append(queue, pair{i: i, j: j - 1})
						grid[i][j-1] = '0'
					}
					if j+1 < column && grid[i][j+1] == '1' {
						queue = append(queue, pair{i: i, j: j + 1})
						grid[i][j+1] = '0'
					}
				}
			}
		}
	}

	return result
}

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands1(grid))
	grid = [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands2(grid))
}
