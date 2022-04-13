package main

import "fmt"

/*
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

示例 1：
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

示例 2：
输入：nums = [0,1]
输出：[[0,1],[1,0]]

示例 3：
输入：nums = [1]
输出：[[1]]
*/

func permute(nums []int) [][]int {
	var ans [][]int
	var backTrack func(nums []int, length int, path []int)
	backTrack = func(nums []int, length int, path []int) {
		if len(nums) == 0 {
			//重新起一个地址存放，防止path上的值被覆盖
			p := make([]int, len(path))
			copy(p, path) //path -> p
			ans = append(ans, p)
		}
		for i := 0; i < length; i++ {
			cur := nums[i]

			path = append(path, cur)
			nums = append(nums[:i], nums[i+1:]...) //把位置i处的元素去除

			backTrack(nums, len(nums), path)

			//撤销之前的操作
			nums = append(nums[:i], append([]int{cur}, nums[i:]...)...)
			path = path[:len(path)-1]
		}
	}
	backTrack(nums, len(nums), []int{})
	return ans
}

func main() {
	nums := []int{5, 4, 6, 2}
	fmt.Println(permute(nums))
}
