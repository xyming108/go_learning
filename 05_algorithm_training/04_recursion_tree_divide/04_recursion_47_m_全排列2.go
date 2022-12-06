package main

import "fmt"

/*
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

示例 1：
输入：nums = [1,1,2]
输出：
[[1,1,2],
[1,2,1],
[2,1,1]]

示例 2：
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

提示：
1 <= nums.length <= 8
-10 <= nums[i] <= 10
nums 中的整数可以相同
*/

func permuteUnique(nums []int) [][]int {
	var ans [][]int
	var backTrack func(nums []int, length int, path []int)

	backTrack = func(nums []int, length int, path []int) {
		if len(nums) == 0 {
			//重新起一个地址存放，防止path上的值被覆盖
			p := make([]int, len(path))
			copy(p, path) //path -> p
			ans = append(ans, p)
			return
		}
		visited := [21]int{}
		for i := 0; i < length; i++ {
			cur := nums[i]
			if visited[nums[i]+10] == 1 { //如果已经存在过了，则剪枝
				continue
			}
			path = append(path, cur)
			visited[nums[i]+10] = 1                //每个元素使用后都标记一下
			nums = append(nums[:i], nums[i+1:]...) //把位置i处的元素去除

			backTrack(nums, len(nums), path)

			//回溯之前的操作，也叫做撤销之前的操作
			nums = append(nums[:i], append([]int{cur}, nums[i:]...)...)
			path = path[:len(path)-1]
		}
	}
	backTrack(nums, len(nums), []int{})
	return ans
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(permuteUnique(nums))
}
