package main

import "fmt"

/*
给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。
请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数字，并以数组的形式返回结果。
你能在不使用额外空间且时间复杂度为 O(n) 的情况下解决这个问题吗?

示例 1：
输入：nums = [4,3,2,7,8,2,3,1]
输出：[5,6]

示例 2：
输入：nums = [1,1]
输出：[2]
*/

//原地修改法
//时间复杂度：O(n)
//空间复杂度：O(1)
/*
遍历nums，每遇到一个数 x，就让nums[x−1] 增加 n。
由于nums中所有数均在[1,n]中，增加以后，这些数必然大于 n。
最后遍历nums，若nums[i]未大于 n，
就说明没有遇到过数i+1，这样就找到了缺失的数字
*/
func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	for _, v := range nums {
		v = (v - 1) % n
		nums[v] += n
	}
	var ans []int
	for i, v := range nums {
		if v <= n {
			ans = append(ans, i+1)
		}
	}
	return ans
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(nums))
}
