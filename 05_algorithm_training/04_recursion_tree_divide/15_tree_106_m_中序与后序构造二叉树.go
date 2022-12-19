package main

/*
给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历，
postorder 是同一棵树的后序遍历，请你构造并返回这颗 二叉树 。

示例 1:
输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
输出：[3,9,20,null,null,15,7]

示例 2:
输入：inorder = [-1], postorder = [-1]
输出：[-1]
*/

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

/*
递归
时间复杂度：O(n)
空间复杂度：O(n)
*/
func buildTree2(inorder []int, postorder []int) *TreeNode {
	indexMap := make(map[int]int, 0)
	for i, v := range inorder {
		indexMap[v] = i
	}
	var buildTrees func(int, int) *TreeNode
	buildTrees = func(inorderLeft int, inorderRight int) *TreeNode {
		if inorderLeft > inorderRight {
			return nil
		}
		//从后序遍历找到当前树的根结点，该结点可以把中序遍历划分成左右两颗子树
		inorderNode := postorder[len(postorder)-1]
		postorder = postorder[:len(postorder)-1]
		root := &TreeNode{
			Val: inorderNode,
		}
		inorderNodeIndex := indexMap[inorderNode]
		root.Right = buildTrees(inorderNodeIndex+1, inorderRight) //先 遍历右子树
		root.Left = buildTrees(inorderLeft, inorderNodeIndex-1)   //后 遍历左子树
		return root
	}
	return buildTrees(0, len(inorder)-1)
}

func main() {

}
