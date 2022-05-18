package main

import "fmt"

/*
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
请你设计并实现时间复杂度为O(n) 的算法解决此问题。

示例 1：
输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。

示例 2：
输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9
由示例2可知重复的算同一个
*/

/*
哈希表
时间复杂度：O(n)
空间复杂度：O(n)
*/
func longestConsecutive(nums []int) int {
	mp := make(map[int]bool)
	for _, v := range nums {
		mp[v] = true
	}
	var result int
	for k, _ := range mp {
		//最坏情况时间复杂度为O(n²)，在这里进行优化，
		//即：只有当一个数是连续序列的第一个数的情况下才会进入内层循环
		//如果一个k数前一个存在，则k的前一个数和k会组成连续的数，
		//则会进入if的内循环for，则会多出很多不必要的枚举
		//最终导致最坏情况时间复杂度为O(n²)
		if !mp[k-1] {
			curV := k
			curLen := 1
			for mp[curV+1] {
				curV++
				curLen++
			}
			if curLen > result {
				result = curLen
			}
		}
	}

	return result
}

func main() {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive(nums))
}
