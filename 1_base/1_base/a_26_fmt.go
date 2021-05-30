package main

import "fmt"

func main() {
	//输入
	var s1 string
	var s2 string
	var s3 string

	fmt.Scan(&s1)
	fmt.Println("s1：", s1)

	fmt.Scanf("%s\n", &s2)
	fmt.Println("s2：", s2)

	fmt.Scanln(&s3)
	fmt.Println("s3：", s3)
	fmt.Println()

	//*******************************************

	var aq = 1
	//输出
	fmt.Print(aq, "不换行")
	fmt.Printf("%d\n", aq)
	fmt.Println(aq)

}
