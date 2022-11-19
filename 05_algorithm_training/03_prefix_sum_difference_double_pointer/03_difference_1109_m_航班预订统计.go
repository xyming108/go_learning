package main

import "fmt"

/*
这里有 n 个航班，它们分别从 1 到 n 进行编号。
有一份航班预订表 bookings ，表中第 i 条预订记录 bookings[i] = [firsti, lasti, seatsi]
意味着在从 firsti 到 lasti （包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi 个座位。
请你返回一个长度为 n 的数组 answer，里面的元素是每个航班预定的座位总数。

示例 1：
输入：bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
输出：[10,55,45,25,25]
解释：
航班编号        1   2   3   4   5
预订记录 1 ：   10  10
预订记录 2 ：       20  20
预订记录 3 ：       25  25  25  25
总座位数：      10  55  45  25  25
因此，answer = [10,55,45,25,25]

示例 2：
输入：bookings = [[1,2,10],[2,2,15]], n = 2
输出：[10,25]
解释：
航班编号        1   2
预订记录 1 ：   10  10
预订记录 2 ：       15
总座位数：      10  25
因此，answer = [10,25]

提示：
1 <= n <= 2 * 10^4
1 <= bookings.length <= 2 * 104
bookings[i].length == 3
1 <= firsti <= lasti <= n
1 <= seatsi <= 10^4
*/

/*
方法：差分
差分其实是前缀和的概念，题目要求对区间[l,r]内的值增量叠加，暴力的做法是直接遍历，
但是这样会增加n的复杂度，差分的方法就是对 >=l 的值全部增加增量v，因为只需要l-r之间，
所有 r 右边的增量是多余的，所以还需要对 >r 的值减去刚才加上的增量，最终就只是对[l,r]之间
进行增量叠加，整个过程只需对l个r进行上述两个操作，相比于遍历整个 r-l 的范围来说，降了一个量级的时间复杂度。

时间复杂度：O(m+n)
空间复杂度：O(1)
*/
func corpFlightBookings(bookings [][]int, n int) []int {
	sums := make([]int, n)
	// 因为题目传进来的值是从1开始的，而我们数组是从0开始的，
	// 所以要注意开闭区间
	for _, v := range bookings {
		sums[v[0]-1] += v[2] //对 v[0]-1 右侧(闭区间)全部增加
		if v[1] < n {        //如果是最后一个下标，不要进行减少，不然会数组越界
			sums[v[1]] -= v[2] //对 v[1] 右侧(开区间)全部减少
		}
	}
	for i := 1; i < n; i++ {
		sums[i] += sums[i-1] //采用前缀和思想，直接相加
	}
	return sums
}

func main() {
	bookings := [][]int{
		{1, 2, 10},
		{2, 3, 20},
		{2, 5, 25},
	}
	n := 5
	fmt.Println(corpFlightBookings(bookings, n))
}
