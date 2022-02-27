package main

import "fmt"

/*
给一个单链表的头节点 head ，请你反转链表，并返回反转后的链表

例如：
输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]

输入：head = []
输出：[]
*/

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

//迭代法
//时间复杂度：O(n)
//空间复杂度：O(1)
func reverseList1(head *ListNode) *ListNode {
	var nPrev *ListNode
	nNow := head
	for nNow != nil {
		tempNode := nNow.Next
		nNow.Next = nPrev
		nPrev = nNow
		nNow = tempNode
	}
	return nPrev
}

//递归法
//时间复杂度：O(n)
//空间复杂度：O(n)
//递归过程
/*
  reverseList: head=1
    reverseList: head=2
	    reverseList: head=3
		    reverseList:head=4
			    reverseList:head=5
					终止返回
				cur = 5
				4.next.next->4，即5->4
			cur=5
			3.next.next->3，即4->3
		cur = 5
		2.next.next->2，即3->2
	cur = 5
	1.next.next->1，即2->1

	最后返回cur
*/
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
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
	//a := reverseList1(node1)
	//for a != nil {
	//	fmt.Println(a.Val)
	//	a = a.Next
	//}

	b := reverseList1(node1)
	for b != nil {
		fmt.Println(b.Val)
		b = b.Next
	}
}
