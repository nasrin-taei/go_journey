package main

import "fmt"

func main() {
	newPrint(1, 3, 10)

	a := []int{5, 8, 10}
	newPrint(a...)

	a = append(a, 2, 3, 4, 5, 7)
	newPrint(a...)

	b := []int{11, 17}
	u := make([]int, len(a))
	copy(u, a)

	//1
	for _, n := range b {
		u = append(u, n)
	}
	newPrint(u...)

	//2
	u2 := append(a, b...)
	newPrint(u2...)
}

func newPrint(p ...int) {
	fmt.Println(p)
}
