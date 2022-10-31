package main

import "fmt"

//查找表法
//用map记录已经遍历过的数值和下标
//空间换时间
//时间复杂度：O(N)
//空间复杂度：O(N)
func twoSum(nums []int, target int) []int {
	mapTemp := make(map[int]int, len(nums))
	for i, v := range nums {
		if j, ok := mapTemp[target-v]; ok {
			return []int{j, i}
		}
		mapTemp[v] = i
	}
	return nil
}

func main() {
	nums := []int{2, 11, 7, 15}
	target := 9

	result1 := twoSum(nums, target)
	fmt.Println(result1)
}
