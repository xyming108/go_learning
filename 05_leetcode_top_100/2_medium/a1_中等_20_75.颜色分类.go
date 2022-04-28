package main

import "fmt"

/*
给定一个包含红色、白色和蓝色、共n 个元素的数组nums，原地对它们进行排序，
使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
我们使用整数 0、1 和 2 分别表示红色、白色和蓝色。
必须在不使用库的sort函数的情况下解决这个问题。

示例 1：
输入：nums = [2,0,2,1,1,0]
输出：[0,0,1,1,2,2]

示例 2：
输入：nums = [2,0,1]
输出：[0,1,2]
*/

/*
单指针
时间复杂度：O(n)
空间复杂度：O(1)
*/
func sortColors1(nums []int) {
	var swap func(nums []int, target int) (targetIndex int)
	swap = func(nums []int, target int) (targetIndex int) {
		for i, v := range nums {
			if v == target {
				nums[i], nums[targetIndex] = nums[targetIndex], nums[i]
				targetIndex++
			}
		}
		return
	}

	index := swap(nums, 0)
	swap(nums[index:], 1)
}

/*
双指针
时间复杂度：O(n)
空间复杂度：O(1)
*/
func sortColors2(nums []int) {
	p0, p1 := 0, 0
	for i, v := range nums {
		if v == 0 {
			nums[p0], nums[i] = nums[i], nums[p0]
			if p0 < p1 {
				nums[p1], nums[i] = nums[i], nums[p1]
			}
			p0++
			p1++
		} else if v == 1 {
			nums[p1], nums[i] = nums[i], nums[p1]
			p1++
		}
	}
}

/*
双指针
时间复杂度：O(n)
空间复杂度：O(1)
*/
func sortColors3(nums []int) {
	p0, p2 := 0, len(nums)-1
	for i := 0; i <= p2; i++ {
		for ; i <= p2 && nums[i] == 2; p2-- {
			nums[p2], nums[i] = nums[i], nums[p2]
		}
		if nums[i] == 0 {
			nums[p0], nums[i] = nums[i], nums[p0]
			p0++
		}
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	//sortColors1(nums)
	//sortColors2(nums)
	sortColors3(nums)
	fmt.Println(nums)
}
