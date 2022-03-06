package main

import "fmt"

/*
给定一个 row x col 的二维网格地图 grid ，其中：grid[i][j] = 1 表示陆地， grid[i][j] = 0 表示水域。
网格中的格子 水平和垂直 方向相连（对角线方向不相连）。整个网格被水完全包围，
但其中恰好有一个岛屿（或者说，一个或多个表示陆地的格子相连组成的岛屿）。

岛屿中没有“湖”（“湖” 指水域在岛屿内部且不和岛屿周围的水相连）。
格子是边长为 1 的正方形。网格为长方形，且宽度和高度均不超过 100 。计算这个岛屿的周长

示例 1：
输入：grid = [[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]
输出：16
解释：它的周长是上面图片中的 16 个黄色的边

示例 2：
输入：grid = [[1]]
输出：4

示例 3：
输入：grid = [[1,0]]
输出：4

leetcode：https://leetcode-cn.com/problems/island-perimeter/
*/

//迭代
//对于一个陆地格子的每条边，它被算作岛屿的周长当且仅当这条边为网格的边界或者相邻的另一个格子为水域。
//因此，我们可以遍历每个陆地格子，看其四个方向是否为边界或者水域，
//如果是，将这条边的 1 加入答案 ans 中即可
//时间复杂度：O(nm)
//空间复杂度：O(1)
func islandPerimeter1(grid [][]int) int {
	type pair struct {
		x, y int
	}
	dir := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n, ans := len(grid), len(grid[0]), 0
	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				for _, d := range dir {
					if x, y := i+d.x, j+d.y; x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == 0 {
						ans++
					}
				}
			}
		}
	}
	return ans
}

//深度优先搜索
//设定值为 22 的格子为已经遍历过的陆地格子
//时间复杂度：O(nm)
//空间复杂度：O(nm)
func islandPerimeter2(grid [][]int) int {
	type pair struct {
		x, y int
	}
	dir := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n, ans := len(grid), len(grid[0]), 0

	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == 0 {
			ans++
			return
		}
		if grid[x][y] == 2 {
			return
		}
		grid[x][y] = 2
		for _, d := range dir {
			dfs(x+d.x, y+d.y)
		}
	}

	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				dfs(i, j)
			}
		}
	}
	return ans
}

func main() {
	grid := [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}, {1, 1, 0, 0}}
	result := islandPerimeter1(grid)
	fmt.Println(result)
	result = islandPerimeter2(grid)
	fmt.Println(result)
}
