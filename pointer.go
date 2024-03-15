package main

import "fmt"

func main() {

	var a = 2
	fmt.Println(a)

	var b *int
	b = &a
	fmt.Println(&a)
	fmt.Println(b)

	a = 7
	fmt.Println(a)

	*b = 8
	fmt.Println(a)
	fmt.Println(*b)

	c, d := 10, 12
	fmt.Printf("c %d, d %d\n", c, d)
	testPointer(c, &d)
	fmt.Printf("c %d, d %d\n", c, d)
}

//1st parameter is pass by value
//2nd parameter is pass by reference
func testPointer(a int, b *int) {
	a += 2
	*b += 2

	fmt.Println("a, b", a, *b)
}
