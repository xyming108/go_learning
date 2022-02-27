package main

/*
给定一棵二叉树，你需要计算它的直径长度。
一棵二叉树的直径长度是任意两个结点路径长度中的最大值。
这条路径可能穿过也可能不穿过根结点。

示例 :
给定二叉树
          1
         / \
        2   3
       / \
      4   5
返回3, 它的长度是路径 [4,2,1,3] 或者[5,2,1,3]。

注意：两结点之间的路径长度是以它们之间边的数目表示。
*/

/*type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}*/

//深度优先搜索
//时间复杂度：O(N)
//空间复杂度：O(Height)
func diameterOfBinaryTree(root *TreeNode) int {
	var ans = 1
	var depth func(root *TreeNode) int
	depth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := depth(root.Left)
		r := depth(root.Right)
		ans = max(l+r+1, ans)
		return max(l, r) + 1
	}
	depth(root)
	return ans - 1
}

func main() {
	node := new(TreeNode)

	diameterOfBinaryTree(node)
}
