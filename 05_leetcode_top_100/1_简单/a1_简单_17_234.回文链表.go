package main

import "fmt"

/*
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true，否则，返回 false
能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题

示例1：
输入：head = [1,2,2,1]
输出：true

输入：head = [1,2]
输出：false
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

//快慢指针法
//时间复杂度：O(n)
//空间复杂度：O(1)
/*
整个流程可以分为以下五个步骤：

找到前半部分链表的尾节点
反转后半部分链表
判断是否回文
恢复链表
返回结果
*/
func isPalindrome1(head *ListNode) bool {
	var reverseList func(head *ListNode) *ListNode
	reverseList = func(head *ListNode) *ListNode {
		var prev, cur *ListNode = nil, head
		for cur != nil {
			nextTmp := cur.Next
			cur.Next = prev
			prev = cur
			cur = nextTmp
		}
		return prev
	}

	var endOfFirstHalf func(head *ListNode) *ListNode
	endOfFirstHalf = func(head *ListNode) *ListNode {
		fast := head
		slow := head
		for fast.Next != nil && fast.Next.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
		}
		return slow
	}

	if head == nil {
		return true
	}

	firstHalfEnd := endOfFirstHalf(head)
	secondHalfStart := reverseList(firstHalfEnd.Next)

	p1 := head
	p2 := secondHalfStart
	result := true
	for result && p2 != nil {
		if p1.Val != p2.Val {
			result = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	firstHalfEnd.Next = reverseList(secondHalfStart)
	return result
}

//将值复制到数组中后用双指针法
//时间复杂度：O(n)
//空间复杂度：O(n)
func isPalindrome2(head *ListNode) bool {
	var vals []int
	for ; head != nil; head = head.Next {
		vals = append(vals, head.Val)
	}
	n := len(vals)
	for i, v := range vals[:n/2] {
		if v != vals[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	node1 := new(ListNode)
	node2 := new(ListNode)
	node3 := new(ListNode)
	node4 := new(ListNode)
	node1.Val = 1
	node1.Next = node2
	node2.Val = 2
	node2.Next = node3
	node3.Val = 2
	node3.Next = node4
	node4.Val = 1
	fmt.Println(isPalindrome1(node1))
	fmt.Println(isPalindrome2(node1))
}
