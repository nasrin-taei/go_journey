package main

import "fmt"

type AInterface interface {
	doA1Thing()
}

type A1Implementation struct {
	s string
}

func (a A1Implementation) doA1Thing() {
	fmt.Println("dddddddddddddddd")
}

func main() {
	var x AInterface
	x = A1Implementation{s: "WHATEVER"}
	x.doA1Thing()

	//z, ok := x.(A1Implementation) //check if the conversion is OK
	z := x.(A1Implementation)
	fmt.Println(z.s)

	var x2 AInterface = z
	x2.doA1Thing()
}
