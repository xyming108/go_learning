package main

import "fmt"

/*
给定一个二叉树，找出其最小深度。
最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
说明：叶子节点是指没有子节点的节点。

示例 1：
输入：root = [3,9,20,null,null,15,7]
	3
   / \
  9  20
    /  \
   15   7
输出：2

示例 2：
输入：root = [2,null,3,null,4,null,5,null,6]
输出：5
*/

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
深度优先
时间复杂度：O(n)
空间复杂度：O(height)
*/
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

/*
广度优先搜索
时间复杂度：O(n)
空间复杂度：最坏情况O(n)
*/
func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue, res := make([]*TreeNode, 0), 0
	queue = append(queue, root)
	for len(queue) > 0 {
		res++
		tempQueue := make([]*TreeNode, 0)
		for _, v := range queue {
			if v.Left == nil && v.Right == nil {
				return res
			}
			if v.Left != nil {
				tempQueue = append(tempQueue, v.Left)
			}
			if v.Right != nil {
				tempQueue = append(tempQueue, v.Right)
			}
		}
		queue = tempQueue
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

	fmt.Println(minDepth(root))
	fmt.Println(minDepth1(root))
}
