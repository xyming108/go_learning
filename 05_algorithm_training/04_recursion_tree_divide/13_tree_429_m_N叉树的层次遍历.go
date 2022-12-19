package main

/*
给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。
树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。

示例 1：
输入：root = [1,null,3,2,4,null,5,6]
输出：[[1],[3,2,4],[5,6]]
*/

type Node struct {
	Val      int
	Children []*Node
}

/*
广度优先搜索
时间复杂度：O(n)
空间复杂度：O(n)
*/
func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	ans := make([][]int, 0)
	levelQueue := make([]*Node, 0)
	levelQueue = append(levelQueue, root)
	for len(levelQueue) > 0 {
		var tmpQueue []*Node
		var tmpVal []int
		for _, v := range levelQueue {
			tmpVal = append(tmpVal, v.Val)
			if len(v.Children) > 0 {
				tmpQueue = append(tmpQueue, v.Children...)
			}
		}
		levelQueue = tmpQueue
		ans = append(ans, tmpVal)
	}
	return ans
}

func main() {

}
