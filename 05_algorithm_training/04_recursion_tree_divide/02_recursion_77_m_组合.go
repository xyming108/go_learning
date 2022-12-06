package main

import "fmt"

/*
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
你可以按 任何顺序 返回答案。

示例 1：
输入：n = 4, k = 2
输出：
[[2,4],
[3,4],
[2,3],
[1,2],
[1,3],
[1,4]]

示例 2：
输入：n = 1, k = 1
输出：[[1]]

提示：
1 <= n <= 20
1 <= k <= n
*/

/*
递归，回溯
时间复杂度：O(C(n, k)*k)，其中C(n, k)表示每次循环中包含的排列组合结果
空间复杂度：O(n+k)
*/
func combine(n int, k int) [][]int {
	res, temp := make([][]int, 0), make([]int, 0)
	var recursion func(cur int)
	recursion = func(cur int) {
		//剪枝：去除多余循环，只有 >= k的个数时才存在答案
		if len(temp)+n-cur+1 < k {
			return
		}

		if len(temp) == k { //匹配到了直接放入结果
			nums := make([]int, k)
			copy(nums, temp)
			res = append(res, nums)
			return
		}

		temp = append(temp, cur) //考虑该值
		recursion(cur + 1)
		temp = temp[:len(temp)-1] //不考虑该值，回溯
		recursion(cur + 1)
	}
	recursion(1)
	return res
}

func main() {
	fmt.Println(combine(5, 3))
}
