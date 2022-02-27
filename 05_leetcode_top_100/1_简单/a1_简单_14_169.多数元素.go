package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

/*
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于⌊ n/2 ⌋的元素。
可以假设数组是非空的，并且给定的数组总是存在多数元素。

设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题

示例1：
输入：[3,2,3]
输出：3

示例2：
输入：[2,2,1,1,1,2,2]
输出：2
*/

//哈希表法
//元素作为key，个数作为value，根据value最大的那个选择key对应的元素即是结果
//时间复杂度：O(n)
//空间复杂度：O(n)
func majorityElement1(nums []int) int {
	tempMap := map[int]int{}
	max := math.MinInt
	final := 0
	for _, v := range nums {
		if _, ok := tempMap[v]; ok {
			tempMap[v]++
		} else {
			tempMap[v] = 1
		}
	}

	for k, v := range tempMap {
		if v > max {
			max = v
			final = k
		}
	}
	return final
}

//排序法
//如果将数组 nums 中的所有元素按照单调递增或单调递减的顺序排序，
//那么下标为 ⌊ n/2 ⌋ 的元素（下标从 0 开始）一定是众数
//时间复杂度：O(nlogn)，将数组排序的时间复杂度为 O(nlogn)
//空间复杂度：O(logn)
func majorityElement2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

//随机化法
//因为超过 ⌊n/2⌋ 的数组下标被众数占据了，这样随机挑选一个下标对应的元素并验证，有很大的概率能找到众数
//时间复杂度：理论上最坏情况下的时间复杂度为 O(∞)，期望的时间复杂度为 O(n)
//空间复杂度：O(1)
func majorityElement3(nums []int) int {
	for {
		randIndex := rand.Intn(len(nums))
		candidate := nums[randIndex]
		count := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] == candidate {
				count++
				if count > len(nums)/2 {
					return nums[i]
				}
			}
		}
	}
}

//投票算法
//维护一个候选众数 candidate 和它出现的次数 count
//遍历数组 nums 中的所有元素，对于每个元素 x，在判断 x 之前，如果 count 的值为 0，先将 x 的值赋予 candidate，随后判断 x：
//如果 x 与 candidate 相等，那么计数器 count 的值增加 1；
//如果 x 与 candidate 不等，那么计数器 count 的值减少 1。
//在遍历完成后，candidate 即为整个数组的众数

//时间复杂度：O(n)
//空间复杂度：O(1)
func majorityElement4(nums []int) int {
	count := 0
	candidate := 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
		}
		if candidate == nums[i] {
			count++
		} else {
			count--
		}
	}
	return candidate
}

func main() {
	s := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement1(s))
	fmt.Println(majorityElement2(s))
	fmt.Println(majorityElement3(s))
	fmt.Println(majorityElement4(s))
}
