package main

import "fmt"

/*
给定一个二叉树的根节点 root ，返回它的 中序 遍历。

示例 1：
       1
        \
         2
        /
       3
输入：root = [1,null,2,3]
输出：[1,3,2]

示例 2：
输入：root = []
输出：[]

示例 3：
输入：root = [1]
输出：[1]

示例 4：
       1
        \
         2
输入：root = [1,null,2]
输出：[1,2]
*/

/*type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}*/

//递归法
//时间复杂度：O(n)，其中 n 为二叉树节点的个数。二叉树的遍历中每个节点会被访问一次且只会被访问一次。
//空间复杂度：O(n)，空间复杂度取决于递归的栈深度，而栈深度在二叉树为一条链的情况下会达到 O(n) 的级别。
func inorderTraversal1(root *TreeNode) (res []int) {
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node != nil {
			inorder(node.Left)
			res = append(res, node.Val)
			inorder(node.Right)
		}
		return
	}
	inorder(root)
	return
}

//迭代法
//时间复杂度：O(n)，其中 n 为二叉树节点的个数。二叉树的遍历中每个节点会被访问一次且只会被访问一次。
//空间复杂度：O(n)，空间复杂度取决于栈深度，而栈深度在二叉树为一条链的情况下会达到 O(n) 的级别。
func inorderTraversal2(root *TreeNode) (res []int) {
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
}

func main() {
	var root = new(TreeNode)
	var root1 = new(TreeNode)
	var root2 = new(TreeNode)

	root.Val = 1
	root.Left = nil
	root.Right = root1

	root1.Val = 2
	root1.Right = nil
	root1.Left = root2

	root2.Val = 3
	root2.Left = nil
	root2.Right = nil

	a := inorderTraversal1(root)
	fmt.Println(a)

	b := inorderTraversal2(root)
	fmt.Println(b)

}
