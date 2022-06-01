package main

func main() {
	s := make([]int, 1000, 1000)
	for index, _ := range s {
		s[index] = index
	}
}
