package main

import "fmt"

/*
给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，
返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。
不要使用额外的空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成
如果在删除重复项之后有 k 个元素，那么 nums 的前 k 个元素应该保存最终结果
不需要考虑数组中超出新长度后面的元素
*/

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	tempValue := nums[0]
	vCount := 1
	for _, v := range nums {
		if tempValue != v {
			nums[vCount] = v
			tempValue = v
			vCount++
		}
	}
	return vCount
}

//双指针
func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func main() {
	num := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	num2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(num))
	fmt.Println(removeDuplicates2(num2))
}
