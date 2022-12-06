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

提示：
1 <= nums.length <= 6
-10 <= nums[i] <= 10
nums 中的所有整数 互不相同
*/

/*
时间复杂度：O(n×n!)
空间复杂度：O(n)

回溯算法 =（递归+纯暴力搜索）因为：有的能用回溯做出来就不错了
特点：抽象
解决问题：组合、切割、子集、排列、棋盘（N皇后）
解题方法：动手画图（切忌纯想象），一般回溯问题都可以变形为n叉树形结构

解题模板：
func {
	if 终止条件 {
		收集结果
		return
	}
	for 元素集合 {
		处理结点
		递归
		回溯
	}
	正常return
}

学习视频地址：https://www.bilibili.com/video/BV1cy4y167mM/
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
			return
		}
		for i := 0; i < length; i++ {
			cur := nums[i]

			path = append(path, cur)
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
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}
