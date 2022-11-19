package main

import "fmt"

/*
给你一个整数数组 nums 和一个整数 k。如果某个连续子数组中恰好有 k 个奇数数字，
我们就认为这个子数组是「优美子数组」。
请返回这个数组中 「优美子数组」 的数目。

示例 1：
输入：nums = [1,1,2,1,1], k = 3
输出：2
解释：包含 3 个奇数的子数组是 [1,1,2,1] 和 [1,2,1,1] 。

示例 2：
输入：nums = [2,4,6], k = 1
输出：0
解释：数列中不包含任何奇数，所以不存在优美子数组。

示例 3：
输入：nums = [2,2,2,1,2,2,1,2,2,2], k = 2
输出：16

提示：
1 <= nums.length <= 50000
1 <= nums[i] <= 10^5
1 <= k <= nums.length
*/

/*
前缀和
时间复杂度：O(n)
空间复杂度：O(n)
*/
func numberOfSubarrays(nums []int, k int) int {
	// cnt存放[优美子数组]出现的次数
	// odd记录累计奇数的个数
	// ans存放最终结果符合条件的[优美子数组]
	cnt, odd, ans := make([]int, len(nums)+1), 0, 0
	cnt[0] = 1
	for _, v := range nums {
		// 奇数&1为1，偶数&1为0
		odd += v & 1
		if odd-k >= 0 {
			ans += cnt[odd-k]
		}
		cnt[odd] += 1
	}

	return ans
}

func main() {
	nums := []int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}
	k := 2
	fmt.Println(numberOfSubarrays(nums, k))
}
