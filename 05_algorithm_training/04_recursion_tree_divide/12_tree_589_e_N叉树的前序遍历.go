package main

/*
给定一个 n 叉树的根节点  root ，返回 其节点值的 前序遍历 。
n 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔（请参见示例）。

示例 1：
输入：root = [1,null,3,2,4,null,5,6]
输出：[1,3,5,6,2,4]
*/

//type Node struct {
//	Val      int
//	Children []*Node
//}

/*
递归法
时间复杂度：O(m)
空间复杂度：O(m)
*/
func preorder(root *Node) []int {
	var ans []int
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		for _, v := range node.Children {
			dfs(v)
		}
	}
	dfs(root)
	return ans
}

func main() {

}
