package main

import "fmt"

/*
给定一个m x n 二维字符网格board 和一个字符串单词word 。
如果word 存在于网格中，返回 true ；否则，返回 false 。
单词必须按照字母顺序，通过相邻的单元格内的字母构成，
其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

示例 1：
输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = 'ABCCED'
输出：true

示例 2：
输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = 'SEE'
输出：true

示例 3：
输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = 'ABCB'
输出：false
*/

/*
回溯
时间复杂度：上界为 O(M*N*(3^L))，其中 M,N 为网格的长度与宽度，L 为字符串 word 的长度
空间复杂度：O(M*N)
*/
func exist(board [][]byte, word string) bool {
	type pair struct {
		x, y int
	}
	directions := []pair{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	h, w := len(board), len(board[0])

	vis := make([][]bool, h)
	for i := 0; i < h; i++ {
		vis[i] = make([]bool, w)
	}

	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		if board[i][j] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		vis[i][j] = true
		for m := 0; m < len(directions); m++ {
			newI, newJ := i+directions[m].x, j+directions[m].y
			if (newI >= 0 && newI < h) && (newJ >= 0 && newJ < w) && !vis[newI][newJ] {
				if check(newI, newJ, k+1) {
					return true
				}
			}
		}
		vis[i][j] = false
		return false
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCCED"
	fmt.Println(exist(board, word))
}
