package main

import "fmt"

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
注意：给定 n 是一个正整数。

示例 1：
输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶

示例 2：
输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶
*/

//暴力递归，动态规划
//转移方程：f(x) = f(x-1)+f(x-2)
//时间复杂度：O(2ⁿ)
//空间复杂度：O(n)
func climbStairs1(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairs1(n-1) + climbStairs1(n-2)
}

//记忆化递归，动态规划
//把已经用过的楼梯数保存起来直接返回，减少递归次数
//时间复杂度：O(n)
//空间复杂度：O(n)
func climbStairs2(n int) int {
	memo := make([]int, n+1, n+1)
	return climbStairsMemo(n, memo)
}
func climbStairsMemo(n int, memo []int) int {
	if memo[n] > 0 {
		return memo[n]
	}
	if n == 1 {
		memo[n] = 1
	} else if n == 2 {
		memo[n] = 2
	} else {
		memo[n] = climbStairsMemo(n-1, memo) + climbStairsMemo(n-2, memo)
	}
	return memo[n]
}

//优化空间复杂度后的动态规划
//滚动数组思想
//时间复杂度：O(n)
//空间复杂度：O(1)
func climbStairs3(n int) int {
	p, q, r := 0, 0, 1
	for i := 0; i < n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

func main() {
	n := 10
	result1 := climbStairs1(n)
	fmt.Println(result1)

	result2 := climbStairs2(n)
	fmt.Println(result2)

	result3 := climbStairs3(n)
	fmt.Println(result3)
}
