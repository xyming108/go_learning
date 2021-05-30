package main

/**
Author: xym
Date: 2021/5/23 16:22
Project: GO_Learning
Description:
*/

/*
2.寻找两个正序数组的中位数 困难
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
示例 3：

输入：nums1 = [0,0], nums2 = [0,0]
输出：0.00000
示例 4：

输入：nums1 = [], nums2 = [1]
输出：1.00000
示例 5：

输入：nums1 = [2], nums2 = []
输出：2.00000
提示：

nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-106 <= nums1[i], nums2[i] <= 106
*/

import (
	"fmt"
)

func main() {
	fmt.Println(findMedianSortedArrays1([]int{}, []int{}))                  // OJ 无此类异常 case
	fmt.Println(findMedianSortedArrays1([]int{1, 3, 5}, []int{2, 4, 6}))    // 3.5
	fmt.Println(findMedianSortedArrays1([]int{1, 3, 5, 7}, []int{2, 4, 6})) // 4
	fmt.Println(findMedianSortedArrays2([]int{1, 3, 5}, []int{2, 4, 6}))    // 3.5
}

// 合并两个数组后直接取中位数即可
// 时间复杂度 O(M+N) // not ok
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	res := merge(nums1, nums2)
	n := len(res)
	if n == 0 {
		return -1
	}
	if n%2 == 0 {
		return float64(res[n/2-1]+res[n/2]) / 2 // len(nums)/2 是中位数 pair 的后一个值
	}
	return float64(res[n/2])
}

// 合并两个有序数组
func merge(nums1, nums2 []int) []int {
	n1, n2 := len(nums1), len(nums2)
	i, j := 0, 0
	res := make([]int, 0, n1+n2)
	for i < n1 && j < n2 {
		switch {
		case nums1[i] < nums2[j]:
			res = append(res, nums1[i])
			i++
		case nums1[i] > nums2[j]:
			res = append(res, nums2[j])
			j++
		default:
			res = append(res, nums1[i], nums2[j])
			i++
			j++
		}
	}

	// 处理未合并部分
	if i < n1 {
		res = append(res, nums1[i:]...) // 循环退出时 i 未被处理，需合并 [i:]
	}
	if j < n2 {
		res = append(res, nums2[j:]...)
	}
	return res
}

//**********************************************************
// 类二分查找分而治之的思路
// O(log(M+N)) // ok
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)
	if n1+n2 == 0 {
		return -1
	}
	if (n1+n2)%2 == 0 {
		l := findKth(nums1, nums2, (n1+n2)/2)
		r := findKth(nums1, nums2, (n1+n2)/2+1) // 此处 +1 与上边 len(mergedNums)/2 同理
		return float64(l+r) / 2
	}
	return float64(findKth(nums1, nums2, (n1+n2)/2+1))
}

// 在两个有序数组中查找第 k 大的数
// 第 k 大的数即 nums 中索引为 k-1 的数，在舍弃区域比较值时都要 -1 处理
func findKth(nums1, nums2 []int, k int) int {
	n1, n2 := len(nums1), len(nums2)
	if n1 > n2 {
		n1, n2 = n2, n1
		nums1, nums2 = nums2, nums1 // 为避免数组长度的分类讨论，先做预处理
	}

	if n1 == 0 {
		return nums2[k-1] // bingo
	}

	if k == 1 {
		return min(nums1[0], nums2[0]) // bingo
	}

	k1 := min(k/2, n1) // 避免越界
	k2 := k - k1       // 不能理想的 k/2, k/2 划分

	// fmt.Println(nums1, k1-1, nums2, k2-1)

	switch {
	case nums1[k1-1] < nums2[k2-1]:
		return findKth(nums1[k1:], nums2, k2) // 彻底舍弃区域 1
	case nums1[k1-1] > nums2[k2-1]:
		return findKth(nums1, nums2[k2:], k1) // 彻底舍弃区域 3
	default:
		return nums1[k1-1] // bingo
	}
}
