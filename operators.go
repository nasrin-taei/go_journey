package main

import "fmt"

func main() {

	a, b := 2, 4
	a = a + b
	fmt.Printf("a = %d\n", a)

	a /= b
	fmt.Printf("a = %d\n", a)
	//fmt.Println("a =", a)

	b++
	fmt.Println("b =", b)

	f := 1.1
	f++
	fmt.Println("f =", f)

	fmt.Println("is 'a' bigger than 'b'? ", a > b)

	fmt.Println("is 'a' not equal and bigger than 'b'?", a != b && a > b)

	x, y, z := false, true, false

	fmt.Println("AND operator between x,y,z: ", x && y && z)
	fmt.Println("OR operator between x,y,z: ", x || y || z)

	fmt.Println(3 == 4 && 4 == 3)

	fmt.Println("NOT OR operator between x,y,z: ", !(x || y || z))
}
