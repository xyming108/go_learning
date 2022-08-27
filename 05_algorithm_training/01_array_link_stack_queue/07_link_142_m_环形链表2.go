package main

import "fmt"

/*
给定一个链表的头节点 head，返回链表开始入环的第一个节点。如果链表无环，则返回null。
如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。
为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置(索引从 0 开始)。
如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
不允许修改链表。

示例 1：
输入：head = [3,2,0,-4], pos = 1
输出：返回索引为 1 的链表节点
解释：链表中有一个环，其尾部连接到第二个节点。

示例2：
输入：head = [1,2], pos = 0
输出：返回索引为 0 的链表节点
解释：链表中有一个环，其尾部连接到第一个节点。

示例 3：
输入：head = [1], pos = -1
输出：返回 null
解释：链表中没有环。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

//方法一：哈希表法
//每次到达一个节点，如果该节点已经存在于哈希表中，则说明该链表是环形链表，否则就将该节点加入哈希表中
//时间复杂度：O(n)
//空间复杂度：O(n)
func detectCycle1(head *ListNode) *ListNode {
	tempHead := head
	tempMap := make(map[*ListNode]struct{})
	for tempHead != nil {
		if _, ok := tempMap[tempHead]; ok {
			return tempHead
		}
		tempMap[tempHead] = struct{}{}
		tempHead = tempHead.Next
	}
	return nil
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

	result := detectCycle1(node1)
	fmt.Println(result.Val)
}
