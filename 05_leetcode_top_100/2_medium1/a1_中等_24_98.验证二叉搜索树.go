package main

import (
	"fmt"
	"math"
)

/*
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
有效二叉搜索树定义如下：
节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。

示例 1：
输入：root = [2,1,3]  //按层次遍历生成二叉树
输出：true

示例 2：
输入：root = [5,1,4,null,null,3,6]  //按层次遍历生成二叉树
输出：false
解释：根节点的值是 5 ，但是右子节点的值是 4 。
*/

/*type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}*/

/*
递归
时间复杂度：O(n)
空间复杂度：O(n)
*/
func isValidBST1(root *TreeNode) bool {
	var helper func(root *TreeNode, lower, upper int) bool
	helper = func(root *TreeNode, lower, upper int) bool {
		if root == nil {
			return true
		}
		if root.Val <= lower || root.Val >= upper {
			return false
		}
		return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
	}
	return helper(root, math.MinInt64, math.MaxInt64)
}

/*
中序遍历
时间复杂度：O(n)
空间复杂度：O(n)
*/
func isValidBST2(root *TreeNode) bool {
	var stack []*TreeNode
	inorder := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}

func main() {
	n1, n2, n3, n4, n5 := new(TreeNode), new(TreeNode), new(TreeNode), new(TreeNode), new(TreeNode)
	n1.Val, n2.Val, n3.Val, n4.Val, n5.Val = 5, 1, 4, 3, 6
	n1.Left, n1.Right = n2, n3
	n3.Left, n3.Right = n4, n5
	fmt.Println(isValidBST1(n1))
	fmt.Println(isValidBST2(n1))
}
