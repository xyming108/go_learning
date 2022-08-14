package main

import "fmt"

/*
给你链表的头节点 head ，每k个节点一组进行翻转，请你返回修改后的链表。
k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是k的整数倍，那么请将最后剩余的节点保持原有顺序。
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
你可以设计一个只用 O(1) 额外内存空间的算法解决此问题吗？

例如：
输入：head = [1,2,3,4,5], k = 2
输出：[2,1,4,3,5]

例如：
输入：head = [1,2,3,4,5], k = 3
输出：[3,2,1,4,5]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseNode(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}
	return tail, head
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	pre := hair
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		nex := tail.Next
		head, tail = ReverseNode(head, tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next
	}

	return hair.Next
}

func main() {
	node1 := new(ListNode)
	node2 := new(ListNode)
	node3 := new(ListNode)
	node4 := new(ListNode)
	node5 := new(ListNode)
	node1.Val = 1
	node1.Next = node2
	node2.Val = 2
	node2.Next = node3
	node3.Val = 3
	node3.Next = node4
	node4.Val = 4
	node4.Next = node5
	node5.Val = 5
	node := reverseKGroup(node1, 2)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}
