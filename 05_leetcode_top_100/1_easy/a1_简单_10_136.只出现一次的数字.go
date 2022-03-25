package main

import "fmt"

/*
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:

输入: [2,2,1]
输出: 1
示例2:

输入: [4,1,2,1,2]
输出: 4
*/

//时间复杂度：O(n)
//空间复杂度：O(n)
func singleNumber(nums []int) int {
	repeate := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		_, exist := repeate[nums[i]]
		if exist {
			repeate[nums[i]] = true
		} else {
			repeate[nums[i]] = false
		}
	}
	for i := 0; i < len(nums); i++ {
		if repeate[nums[i]] == false {
			return nums[i]
		}
	}
	return 0
}

//时间复杂度：O(n)
//空间复杂度：O(n)
func singleNumber1(nums []int) int {
	repeate := make(map[int]bool)
	final := 0
	for i := 0; i < len(nums); i++ {
		if _, exist := repeate[nums[i]]; !exist {
			repeate[nums[i]] = true
			final += nums[i]
		} else {
			final -= nums[i]
		}
	}
	return final
}

//异或
//时间复杂度：O(n)
//空间复杂度：O(1)
func singleNumber2(nums []int) int {
	s := 0
	for i:=0; i<len(nums); i++ {
		s = s ^ nums[i]
	}
	return s
}

func main() {
	a := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(a))

	fmt.Println(singleNumber1(a))

	fmt.Println(singleNumber2(a))

}
