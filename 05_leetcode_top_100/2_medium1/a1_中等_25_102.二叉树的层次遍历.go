package main

import "fmt"

/*
给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。

示例 1：
输入：root = [3,9,20,null,null,15,7]
输出：[[3],[9,20],[15,7]]

示例 2：
输入：root = [1]
输出：[[1]]

示例 3：
输入：root = []
输出：[]
*/

/*type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}*/

/*
广度优先搜索
时间复杂度：O(n)
空间复杂度：O(n)
*/
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ret [][]int
	q := []*TreeNode{root} //根结点赋值
	var i int
	for len(q) > 0 {
		ret = append(ret, []int{}) //给ret增加一个初始化
		var p []*TreeNode
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p //把下一层结点赋值给q
		i++   //层数加一
	}

	return ret
}

func main() {
	n1, n2, n3, n4, n5 := new(TreeNode), new(TreeNode), new(TreeNode), new(TreeNode), new(TreeNode)
	n1.Val, n2.Val, n3.Val, n4.Val, n5.Val = 3, 9, 20, 15, 7
	n1.Left, n1.Right = n2, n3
	n3.Left, n3.Right = n4, n5
	fmt.Println(levelOrder(n1))
}
