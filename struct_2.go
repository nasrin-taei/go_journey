package main

import "fmt"

type rectangle struct {
	length, width int
}

func area(coordinate rectangle) int {
	return coordinate.length * coordinate.width
}

func perimeter(coordinate rectangle) int {
	return (coordinate.length + coordinate.width) * 2

}
func main() {

	myrectangle := rectangle{2, 6}
	fmt.Println("for ", myrectangle)
	fmt.Println("the area is", area(myrectangle))
	fmt.Println("the perimeter is", perimeter(myrectangle))
	myrectangle.length *= 2
	myrectangle.width *= 2
	fmt.Println("for", myrectangle)
	fmt.Println("the area is", area(myrectangle))
	fmt.Println("the perimeter is", perimeter(myrectangle))

}
