package main

import "fmt"

/*
给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3

但是下面这个[1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3
*/

/*type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}*/

//递归法
//时间复杂度：O(n)
//空间复杂度：O(n)
/*可以实现这样一个递归函数，通过「同步移动」两个指针的方法来遍历这棵树，
pp 指针和 qq 指针一开始都指向这棵树的根，随后 pp 右移时，qq 左移，pp 左移时，qq 右移。
每次检查当前 pp 和 qq 节点的值是否相等，如果相等再判断左右子树是否对称*/
func isSymmetric1(root *TreeNode) bool {
	return check(root, root)
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}

//迭代法
//时间复杂度：O(n)
//空间复杂度：O(n)
//用队列模拟递归的实现
func isSymmetric2(root *TreeNode) bool {
	u, v := root, root
	var queue []*TreeNode
	queue = append(queue, u)
	queue = append(queue, v)
	for len(queue) > 0 {
		u, v = queue[0], queue[1]
		queue = queue[2:]
		if u == nil && v == nil {
			continue
		}
		if u == nil || v == nil {
			return false
		}
		if u.Val != v.Val {
			return false
		}
		queue = append(queue, u.Left)
		queue = append(queue, v.Right)

		queue = append(queue, u.Right)
		queue = append(queue, v.Left)
	}
	return true
}

func main() {
	var n1 = new(TreeNode)
	var n2 = new(TreeNode)
	var n3 = new(TreeNode)
	var n4 = new(TreeNode)
	var n5 = new(TreeNode)
	var n6 = new(TreeNode)
	var n7 = new(TreeNode)
	n1.Val = 1
	n2.Val = 2
	n3.Val = 2
	n4.Val = 3
	n5.Val = 4
	n6.Val = 4
	n7.Val = 3

	n1.Left = n2
	n1.Right = n3

	n2.Left = n4
	n2.Right = n5

	n3.Left = n6
	n3.Right = n7

	result1 := isSymmetric1(n1)
	fmt.Println(result1)

	result2 := isSymmetric2(n1)
	fmt.Println(result2)
}
