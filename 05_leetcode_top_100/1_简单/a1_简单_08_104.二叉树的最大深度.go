package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明:叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度3
*/

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}


//深度优先搜索
//时间复杂度：O(n)
//空间复杂度：O(height)
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

//广度优先搜索
//时间复杂度：O(n)
//空间复杂度：最坏情况O(n)
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var queue []*TreeNode
	queue = append(queue, root)
	ans := 0
	for len(queue) > 0 {
		sz := len(queue)
		for sz > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			sz--
		}
		ans++
	}
	return ans
}

func main() {
	var root = new(TreeNode)
	var root1 = new(TreeNode)
	var root2 = new(TreeNode)
	var root3 = new(TreeNode)
	var root4 = new(TreeNode)

	root.Val = 3
	root.Left = root1
	root.Right = root2

	root1.Val = 9
	root2.Val = 20

	root2.Left = root3
	root2.Right = root4

	root3.Val = 15
	root4.Val = 7

	fmt.Println(maxDepth(root))
	fmt.Println(maxDepth1(root))
}
