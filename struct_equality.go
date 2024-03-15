package main

import "fmt"

type A struct {
	num  int
	num2 int
}

type B struct {
	num int
}

func main() {
	a := A{num: 1}
	a2 := A{num: 1, num2: 2}

	//b := B{1}
	//if a == b {
	//	fmt.Println("equal")
	//}

	if a == a2 {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}
