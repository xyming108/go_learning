package main

/*
给定一个二叉树的根节点 root，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

示例 1：
输入：root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
输出：3
解释：和等于 8 的路径有 3 条，如图所示。

示例 2：
输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
输出：3
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
深度优先搜索
时间复杂度：O(n^2)
空间复杂度：O(n)
*/
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := rootSum(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}

func rootSum(root *TreeNode, targetSum int) (res int) {
	if root == nil {
		return
	}
	val := root.Val
	if val == targetSum {
		res++
	}
	res += rootSum(root.Left, targetSum-val)
	res += rootSum(root.Right, targetSum-val)
	return
}

/*
前缀和
时间复杂度：O(n)
空间复杂度：O(n)
*/
func pathSum2(root *TreeNode, targetSum int) (ans int) {
	preSum := map[int64]int{0: 1} // 0:1 表示如果 curr-target 结果为0，则算一个结果
	var dfs func(*TreeNode, int64)
	dfs = func(node *TreeNode, curr int64) {
		if node == nil {
			return
		}
		curr += int64(node.Val)
		ans += preSum[curr-int64(targetSum)] //如果 curr-target 结果为0，则算一个结果
		preSum[curr]++                       //往hash表加入当前前缀和
		dfs(node.Left, curr)                 //开始向下深度优先遍历
		dfs(node.Right, curr)
		preSum[curr]-- //回溯，往hash表删除当前前缀和
		return
	}
	dfs(root, 0)
	return
}

func main() {

}
