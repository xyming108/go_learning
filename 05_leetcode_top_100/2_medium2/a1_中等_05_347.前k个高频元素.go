package main

import (
	"container/heap"
	"fmt"
)

/*
给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小

示例 1:
输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]

示例 2:
输入: nums = [1], k = 1
输出: [1]
*/

/*
方法一：大顶堆
时间复杂度：O(nlog(k)), n为数组长度，k为操作堆的时间
空间复杂度：O(N)
*/
func topKFrequent1(nums []int, k int) []int {
	mp := make(map[int]int)
	for _, v := range nums {
		mp[v]++
	}
	h := &IHeap{}
	heap.Init(h)
	for key, value := range mp {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

type IHeap [][2]int

func (I IHeap) Len() int           { return len(I) }
func (I IHeap) Less(i, j int) bool { return I[i][1] < I[j][1] } //大顶堆
func (I IHeap) Swap(i, j int)      { I[i], I[j] = I[j], I[i] }
func (I *IHeap) Push(x any)        { *I = append(*I, x.([2]int)) }
func (I *IHeap) Pop() any {
	old := *I
	n := len(old)
	x := old[n-1]
	*I = old[0 : n-1]
	return x
}

/*
方法二：快速排序
时间复杂度：O(n²)，平均 nlog(n)
空间复杂度：O(n)
https://leetcode.cn/problems/top-k-frequent-elements/solution/qian-k-ge-gao-pin-yuan-su-by-leetcode-solution/
*/
func topKFrequent2(nums []int, k int) []int {
	mp := make(map[int]int)
	for _, v := range nums {
		mp[v]++
	}
	var values [][]int
	for k, v := range mp {
		values = append(values, []int{k, v})
	}

	quickSort(values, 0, len(values)-1)

	var result []int
	for _, v := range values {
		result = append(result, v[0])
	}

	return result[len(result)-k:]
}

func quickSort(a [][]int, left, right int) {
	if left >= right {
		return
	}
	l, r := left, right
	pivot := a[left]
	for left < right {
		for left < right && a[right][1] >= pivot[1] {
			right--
		}
		a[left] = a[right]
		for left < right && a[left][1] <= pivot[1] {
			left++
		}
		a[right] = a[left]
	}
	a[left] = pivot
	middle := left
	quickSort(a, l, middle-1)
	quickSort(a, middle+1, r)
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3, 4, 5}
	k := 2
	fmt.Println(topKFrequent1(nums, k))
	fmt.Println(topKFrequent2(nums, k))
}
