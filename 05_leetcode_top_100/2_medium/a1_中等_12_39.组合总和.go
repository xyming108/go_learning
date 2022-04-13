package main

import "fmt"

/*
给你一个 无重复元素 的整数数组candidates 和一个目标整数target，
找出candidates中可以使数字和为目标数target 的 所有不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
对于给定的输入，保证和为target 的不同组合数少于 150 个。

示例1：
输入：candidates = [2,3,6,7], target = 7
输出：[[2,2,3],[7]]
解释：
2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
7 也是一个候选， 7 = 7 。
仅有这两种组合。

示例2：
输入: candidates = [2,3,5], target = 8
输出: [[2,2,2,2],[2,3,3],[3,5]]

示例 3：
输入: candidates = [2], target = 1
输出: []
*/

/*
搜索回溯，无剪枝
时间复杂度：O(S)
空间复杂度：O(target)
*/
func combinationSum1(candidates []int, target int) (ans [][]int) {
	comb := []int{}
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		// 直接跳过
		dfs(target, idx+1)
		// 选择当前数
		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return
}

func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var dfs func(start int, temp []int, sum int)
	dfs = func(start int, temp []int, sum int) {
		if sum >= target {
			if sum == target {
				newTemp := make([]int, len(temp))
				copy(newTemp, temp)
				res = append(res, newTemp)
			}
			return
		}
		for i := start; i < len(candidates); i++ {
			temp = append(temp, candidates[i])
			dfs(i, temp, sum+candidates[i])
			temp = temp[:len(temp)-1]
		}
	}
	dfs(0, []int{}, 0)
	return res
}

func main() {
	candidates := []int{2, 3, 6, 7}
	fmt.Println(combinationSum1(candidates, 7))
	fmt.Println(combinationSum2(candidates, 7))
}
