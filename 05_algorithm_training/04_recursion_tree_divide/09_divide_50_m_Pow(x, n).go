package main

import "fmt"

/*
实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，xn ）。

示例 1：
输入：x = 2.00000, n = 10
输出：1024.00000

示例 2：
输入：x = 2.10000, n = 3
输出：9.26100

示例 3：
输入：x = 2.00000, n = -2
输出：0.25000
解释：2-2 = 1/22 = 1/4 = 0.25

提示：
-100.0 < x < 100.0
-231 <= n <= 231-1
-104 <= xn <= 104
*/

/*
快速幂（分治）+递归
时间复杂度：log(n)
空间复杂度：log(n)
*/
func myPow(x float64, n int) float64 {
	var recursion func(float64, int) float64
	recursion = func(x float64, n int) float64 {
		if n == 0 {
			return 1
		}
		if n%2 == 1 { //n为奇数
			return x * recursion(x, n-1)
		}
		return recursion(x*x, n/2) //n为偶数
	}

	if n >= 0 { // n为正数
		return recursion(x, n)
	}
	return 1 / recursion(x, -n) //n为负数
}

func main() {
	x := 2.10000
	n := 3
	fmt.Println(myPow(x, n))
}
