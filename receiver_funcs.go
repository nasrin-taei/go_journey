package main

import "fmt"

type Coordinate struct {
	x, y int
}

func shiftBy(x, y int, coord *Coordinate) {
	coord.x += x
	coord.y += y
}

func (rc *Coordinate) shiftBy(x, y int) {
	rc.x += x
	rc.y += y
}

func (rc Coordinate) printInfo() {
	fmt.Println("X is", rc.x, "and Y is", rc.y)
}

func main() {
	c := Coordinate{5, 5}
	shiftBy(1, 1, &c)
	c.shiftBy(1, 1)
	c.printInfo()
}
