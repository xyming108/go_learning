package main

import "fmt"

//golang中最大值和最小值
func maxmin() {
	//math包

	//最大值
	const UINT_MAX = ^uint(0) //18446744073709551615
	//最小值
	const UINT_MIN = 0
	fmt.Println(UINT_MIN, UINT_MAX)

	//根据补码，其最大值二进制表示，首位0，其余1
	const INT_MAX = int(^uint(0) >> 1) //-9223372036854775808
	//根据补码，其最小值二进制表示，首位1，其余0
	const INT_MIN = ^INT_MAX //9223372036854775807
	fmt.Println(INT_MIN, INT_MAX)

	//无穷大，无穷小
	var zero float64
	inf := 1 / zero
	n_inf := -1 / zero
	fmt.Println(inf, n_inf)
	fmt.Println(99999999999999999 < inf)
	fmt.Println(99999999999999999 > n_inf)
	fmt.Printf("%T\n", inf)
	fmt.Printf("%T\n", n_inf)
}

//golang中swap
func swap() {
	nums := []int{3, 5}
	nums[0], nums[1] = nums[1], nums[0]
	fmt.Println(nums)
}

func main() {
	fmt.Println("------------------------------------最大值最小值")
	maxmin()
	fmt.Println("------------------------------------交换")
	swap()
	fmt.Println("------------------------------------")
}
