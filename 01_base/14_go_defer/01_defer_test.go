package main

import (
	"testing"
)

func def1() (a int) {
	a = 1
	defer func() {
		a++
	}()

	return a
}

func def2() (a int) {
	a = 1
	func() {
		a++
	}()
	return a
}

func BenchmarkDefer1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		def1()
	}
}

func BenchmarkDefer2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		def2()
	}
}