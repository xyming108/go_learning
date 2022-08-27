package main

import "fmt"

/*
给一个链表的头节点 head ，判断链表中是否有环。
如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。
为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
注意：pos 不作为参数进行传递。仅仅是为了标识链表的实际情况。
如果链表中存在环，则返回 true 。 否则，返回 false 。

示例：
输入：head = [3,2,0,-4]，pos = 1
输出：true
解释：链表中有一个环，其尾部连接到第二个节点
*/

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

//方法一：哈希表法
//每次到达一个节点，如果该节点已经存在于哈希表中，则说明该链表是环形链表，否则就将该节点加入哈希表中
//时间复杂度：O(n)
//空间复杂度：O(n)
func hasCycle1(head *ListNode) bool {
	tempMap :=  map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := tempMap[head]; ok {
			return true
		}
		tempMap[head] = struct{}{}
		head = head.Next
	}
	return false
}

//方法二：快慢指针法
//一个慢指针每次移动一步，一个快指针一次移动两步，当快指针追上慢指针后说明存在环
//时间复杂度：O(n)
//空间复杂度：O(1)
func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

func main() {
	node1 := new(ListNode)
	node2 := new(ListNode)
	node3 := new(ListNode)
	node4 := new(ListNode)
	node1.Val = 3
	node1.Next = node2
	node2.Val = 2
	node2.Next = node3
	node3.Val = 0
	node3.Next = node4
	node4.Val = -4
	node4.Next = node2

	cycle1 := hasCycle1(node1)
	fmt.Println(cycle1)

	cycle2 := hasCycle2(node1)
	fmt.Println(cycle2)

}
