package main

import "fmt"

/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例1:
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

//递归法
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var result *ListNode
	if l1.Val <= l2.Val {
		result = l1
		result.Next = mergeTwoLists1(l1.Next, l2)
	} else {
		result = l2
		result.Next = mergeTwoLists1(l1, l2.Next)
	}

	return result
}

//迭代法
//会比用递归占用更少的空间
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	nodeTemp := &ListNode{}
	current := nodeTemp

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}
	return nodeTemp.Next
}

func main() {
	var h1 = new(ListNode)
	h1.Val = 1
	var h2 = new(ListNode)
	h2.Val = 2
	var h3 = new(ListNode)
	h3.Val = 4
	h1.Next = h2
	h2.Next = h3
	h3.Next = nil

	var h11 = new(ListNode)
	h11.Val = 1
	var h22 = new(ListNode)
	h22.Val = 3
	var h33 = new(ListNode)
	h33.Val = 4
	h11.Next = h22
	h22.Next = h33
	h33.Next = nil

	//result1 := mergeTwoLists1(h1, h11)
	//fmt.Println(result1)

	result2 := mergeTwoLists2(h1, h11)
	fmt.Println(result2)
}
