package main

import "fmt"

/*
给定一个二叉树，找出其最大深度。
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3
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
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

/*
广度优先搜索
时间复杂度：O(n)
空间复杂度：最坏情况O(n)
*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue, res := make([]*TreeNode, 0), 0
	queue = append(queue, root)
	for len(queue) > 0 {
		tempQueue := make([]*TreeNode, 0)
		for _, v := range queue {
			if v.Left != nil {
				tempQueue = append(tempQueue, v.Left)
			}
			if v.Right != nil {
				tempQueue = append(tempQueue, v.Right)
			}
		}
		queue = tempQueue
		res++
	}
	return res
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
