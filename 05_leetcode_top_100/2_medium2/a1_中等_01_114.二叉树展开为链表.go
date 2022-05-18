package main

/*
给你二叉树的根结点 root ，请你将它展开为一个单链表：
展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。
你可以使用原地算法（O(1) 额外空间）展开这棵树吗？

示例 1：
输入：root = [1,2,5,3,4,null,6]
输出：[1,null,2,null,3,null,4,null,5,null,6]

示例 2：
输入：root = []
输出：[]

示例 3：
输入：root = [0]
输出：[0]
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
前序遍历+递归实现
时间复杂度：O(n)
空间复杂度：O(n)
*/
func flatten1(root *TreeNode) {
	var preorderTraversal func(root *TreeNode) []*TreeNode
	preorderTraversal = func(root *TreeNode) []*TreeNode {
		var list []*TreeNode
		if root != nil {
			list = append(list, root)
			list = append(list, preorderTraversal(root.Left)...)
			list = append(list, preorderTraversal(root.Right)...)
		}
		return list
	}

	list := preorderTraversal(root)
	for i := 1; i < len(list); i++ {
		prev, cur := list[i-1], list[i]
		prev.Left, prev.Right = nil, cur
	}
}

/*
前序遍历+迭代实现
时间复杂度：O(n)
空间复杂度：O(n)
*/
func flatten2(root *TreeNode) {
	var list []*TreeNode
	var stack []*TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			list = append(list, node)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		node = node.Right
		stack = stack[0 : len(stack)-1]
	}

	for i := 1; i < len(list); i++ {
		prev, cur := list[i-1], list[i]
		prev.Left, prev.Right = nil, cur
	}
}

func main() {

}
