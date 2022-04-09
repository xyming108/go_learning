package main

import "fmt"

/*
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点

示例 1：
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]

示例 2：
输入：head = [1], n = 1
输出：[]

示例 3：
输入：head = [1,2], n = 1
输出：[1]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
写法一，无头结点
时间复杂度：O(n)
空间复杂度：O(1)
*/
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	var length int
	node := head
	for node != nil {
		node = node.Next
		length++
	}
	//如果删除的是第一个结点，单独判断
	if length == n {
		return head.Next
	}
	index := length - n + 1
	h := head
	c := h
	for i := 1; i <= index; i++ {
		if i != index-1 {
			h = h.Next
		} else {
			h.Next = h.Next.Next
		}
	}
	return c
}

/*
写法二，有头结点
时间复杂度：O(n)
空间复杂度：O(1)
*/
func getLength(head *ListNode) (length int) {
	for ; head != nil; head = head.Next {
		length++
	}
	return
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	length := getLength(head)
	dummy := &ListNode{0, head}
	cur := dummy
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}

func main() {
	n1, n2, n3, n4, n5 := new(ListNode), new(ListNode), new(ListNode), new(ListNode), new(ListNode)
	n1.Val = 1
	n2.Val = 2
	n3.Val = 3
	n4.Val = 4
	n5.Val = 5
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	result := removeNthFromEnd1(n1, 2)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
