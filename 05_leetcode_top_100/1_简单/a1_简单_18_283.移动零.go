package main

import "fmt"

/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
请注意 ，必须在不复制数组的情况下原地对数组进行操作

示例 1:
输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]

示例 2:
输入: nums = [0]
输出: [0]
*/

//双指针，自己的写法
//时间复杂度：O(n)
//空间复杂度：O(1)
func moveZeroes1(nums []int) {
	length := len(nums)
	if length == 0 || length == 1 {
		return
	}
	prev, cur := 0, 1
	for cur < length && prev != cur {
		for nums[prev] != 0 && prev < cur {
			prev++
		}
		if nums[prev] == 0 && nums[cur] == 0 {
			cur++
		}
		if cur < length && nums[prev] == 0 && nums[cur] != 0 {
			tempData := nums[prev]
			nums[prev] = nums[cur]
			nums[cur] = tempData
			prev++
			cur++
		} else {
			cur++
		}
	}
}

//双指针，官方写法
//时间复杂度：O(n)
//空间复杂度：O(1)
func moveZeroes2(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes1(nums)
	fmt.Println(nums)
}
