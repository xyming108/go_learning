package main

import "fmt"

/*
给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。

示例 1：
输入：n = 3
输出：5

示例 2：
输入：n = 1
输出：1

解析：https://leetcode-cn.com/problems/unique-binary-search-trees/solution/bu-tong-de-er-cha-sou-suo-shu-by-leetcode-solution/
*/

/*
动态规划
转移方程：F(i,n)=G(i−1)⋅G(n−i)
G(i−1)表示根结点左边
G(n−i)表示根结点右边
边界情况：G(0)=1,G(1)=1
时间复杂度 : O(n^2)
空间复杂度 : O(n)
*/
func numTrees1(n int) int {
	G := make([]int, n+1)
	G[0], G[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			G[i] += G[j-1] * G[i-j]
		}
	}
	return G[n]
}

/*
数学公式法
时间复杂度 : O(n)
空间复杂度 : O(1)
*/
func numTrees2(n int) int {
	C := 1
	for i := 0; i < n; i++ {
		C = C * 2 * (2*i + 1) / (i + 2)
	}
	return C
}

func main() {
	fmt.Println(numTrees1(3))
	fmt.Println(numTrees2(3))
}
