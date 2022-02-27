package main

func fooss(b int) *int {
	a1 := new(int)
	a2 := new(int)
	a3 := new(int)
	a4 := new(int)
	a5 := new(int)

	for i := 0; i < 5; i++ {
		println(b, a1, a2, a3, a4, a5)
	}

	return a3
}

func main() {
	p := fooss(666)
	println(*p, p)
}
