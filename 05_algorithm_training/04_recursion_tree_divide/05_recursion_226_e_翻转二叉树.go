package main

/*
给一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点

例如：
输入：root = [4,2,7,1,3,6,9]
输出：[4,7,2,9,6,3,1]

输入：root = []
输出：[]
*/

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

//迭代法
//时间复杂度：O(n)
//空间复杂度：O(n)
func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		tempA := queue[0]
		queue = queue[1:]
		tempData := tempA.Left
		tempA.Left = tempA.Right
		tempA.Right = tempData
		if tempA.Left != nil {
			queue = append(queue, tempA.Left)
		}
		if tempA.Right != nil {
			queue = append(queue, tempA.Right)
		}
	}
	return root
}

//递归法
//时间复杂度：O(n)
//空间复杂度：O(log(n))，最坏情况下O(n)
func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree2(root.Left)
	right := invertTree2(root.Right)
	root.Left = right
	root.Right = left
	return root
}

func main() {
	root, l1, l2, r1, r2 := new(TreeNode), new(TreeNode), new(TreeNode), new(TreeNode), new(TreeNode)
	root.Val, root.Left, root.Right = 4, l1, r1
	l1.Val, l1.Left, l1.Right = 2, l2, r2
	r1.Val, l2.Val, r2.Val = 7, 1, 3
	invertTree2(root)
}
