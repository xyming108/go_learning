package main

import "fmt"

/*
给定两个整数数组preorder 和 inorder，其中preorder 是二叉树的先序遍历，
inorder是同一棵树的中序遍历，请构造二叉树并返回其根节点。
preorder 和 inorder 均 无重复 元素.

示例 1:
输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
输出: [3,9,20,null,null,15,7]

示例 2:
输入: preorder = [-1], inorder = [-1]
输出: [-1]
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
递归
时间复杂度：O(n)
空间复杂度：O(n)
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

func main() {
	preorder := []int{3,9,20,15,7}
	inorder := []int{9,3,15,20,7}
	fmt.Println(buildTree(preorder, inorder))
}
