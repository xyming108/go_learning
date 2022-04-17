package main

import (
	"fmt"
	"sort"
)

/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。

示例 1：
输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

示例2：
输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
*/

/*
排序
时间复杂度：O(nlog(n))
空间复杂度：O(log(n))
*/
func merge(intervals [][]int) [][]int {
	var maxNum func(i, j int) int
	maxNum = func(i, j int) int {
		if i > j {
			return i
		} else {
			return j
		}
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := make([][]int, 0, len(intervals))
	for i := 0; i < len(intervals); i++ {
		L, R := intervals[i][0], intervals[i][1]
		if len(result) == 0 || result[len(result)-1][1] < L {
			result = append(result, []int{L, R})
		} else {
			result[len(result)-1][1] = maxNum(result[len(result)-1][1], R)
		}
	}
	return result
}

func main() {
	intervals := [][]int{{1, 4}, {0, 3}, {3, 5}, {2, 9}, {2, 7}}
	fmt.Println(merge(intervals))
}
