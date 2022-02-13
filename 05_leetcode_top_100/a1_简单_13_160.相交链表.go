package main

/*
给两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点，
如果两个链表不存在相交节点，返回 null。
题目保证整个链式结构中不存在环，且函数返回结果后，链表必须保持其原始结构

设计一个时间复杂度 O(m + n) 、仅用 O(1) 内存的解决方案
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

//哈希集合法
//思路：遍历链表headA，并将链表headA中的每个节点加入哈希集合中，
//然后遍历链表headB，对于遍历到的每个节点，判断该节点是否在哈希集合中。
//时间复杂度：O(m+n)
//空间复杂度：O(n)
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	tempMap := map[*ListNode]bool{}
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		tempMap[tmp] = true
	}
	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if tempMap[tmp] {
			return tmp
		}
	}
	return nil
}

//双指针法
//时间复杂度：O(m+n)
//空间复杂度：O(1)
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}

		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}

func main() {

}
