package main

import "sort"

/*
给你两个按 非递减顺序 排列的整数数组nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。

注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，
其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。
*/

// 时间复杂度：O((m+n)log(m+n))。
// 空间复杂度：O(log(m+n))。
func merge1(nums1 []int, m int, nums2 []int, _ int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

// 自己写的 双指针
// 时间复杂度：O(m+n)
// 空间复杂度：O(m+n)
func merge2(nums1 []int, m int, nums2 []int, n int) {
	num0 := make([]int, m, m)
	for i := 0; i < m; i++ {
		num0[i] = nums1[i]
	}
	i, j, k := 0, 0, 0
	for i < m && j < n {
		if num0[i] < nums2[j] {
			nums1[k] = num0[i]
			i++
		} else {
			nums1[k] = nums2[j]
			j++
		}
		k++
	}
	if i == m {
		for p := k; p < m+n; p++ {
			nums1[p] = nums2[j]
			j++
		}
	} else if j == n {
		for p := k; p < m+n; p++ {
			nums1[p] = num0[i]
			i++
		}
	}
}

// 官方 双指针
// 时间复杂度：O(m+n)
// 空间复杂度：O(m+n)
func merge3(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	p1, p2 := 0, 0
	for {
		if p1 == m {
			sorted = append(sorted, nums2[p2:]...)
			break
		}
		if p2 == n {
			sorted = append(sorted, nums1[p1:]...)
			break
		}
		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}
	copy(nums1, sorted)
}

// 官方 逆向双指针
// 时间复杂度：O(m+n)
// 空间复杂度：O(m+n)
func merge4(nums1 []int, m int, nums2 []int, n int) {
	for p1, p2, tail := m-1, n-1, m+n-1; p1 >= 0 || p2 >= 0; tail-- {
		var cur int
		if p1 == -1 {
			cur = nums2[p2]
			p2--
		} else if p2 == -1 {
			cur = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			cur = nums1[p1]
			p1--
		} else {
			cur = nums2[p2]
			p2--
		}
		nums1[tail] = cur
	}
}

func main() {
	num1 := []int{1, 2, 3, 0, 0, 0}
	num2 := []int{2, 5, 6}
	m, n := 3, 3
	merge2(num1, m, num2, n)
}
