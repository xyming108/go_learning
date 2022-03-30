package main

import (
	"fmt"
)

/*
给你两个非空的链表，表示两个非负的整数。
它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0开头。

示例 1：
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.

示例 2：
输入：l1 = [0], l2 = [0]
输出：[0]

示例 3：
输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
模拟法
时间复杂度：O(max(m,n))
空间复杂度：O(1)
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
}

func main() {
	l1 := new(ListNode)
	l11 := new(ListNode)
	l111 := new(ListNode)
	l1.Val = 2
	l1.Next = l11
	l11.Val = 4
	l11.Next = l111
	l111.Val = 3
	l111.Next = nil

	l2 := new(ListNode)
	l22 := new(ListNode)
	l222 := new(ListNode)
	l2.Val = 5
	l2.Next = l22
	l22.Val = 4
	l22.Next = l222
	l222.Val = 6
	l222.Next = nil

	result := addTwoNumbers(l1, l2)
	fmt.Println(result)
}
