package main

func foo(b int) *int {
	a1 := 11
	a2 := 12
	a3 := 13
	a4 := 14
	a5 := 15

	for i := 0; i < 5; i++ {
		println(&b, &a1, &a2, &a3, &a4, &a5)
	}

	return &a3
}

func main() {
	p := foo(666)
	println(*p, p)
}
