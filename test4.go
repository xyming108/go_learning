package main

import "fmt"

type Template interface {
	hello()
}

func saySome(t Template) {
	t.hello()
}

type Go struct {
	Str string
}

func (g *Go) hello() {
	fmt.Println("=====hello go: ", g.Str)
}

type Cpp struct {
	Str string
}

func (c *Cpp) hello() {
	fmt.Println("=====hello c++: ", c.Str)
}

func main() {
	g := &Go{
		Str: "go语言",
	}
	c := &Cpp{
		Str: "C++语言",
	}

	saySome(g)
	saySome(c)
}
