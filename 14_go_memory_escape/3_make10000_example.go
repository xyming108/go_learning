package main

func main() {
	s := make([]int, 10000, 10000)
	for index, _ := range s {
		s[index] = index
	}
}
