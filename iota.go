package main

import "fmt"

type Direction byte

const (
	point_a Direction = iota
	point_b
	point_c
	point_d
)

func (d Direction) getName() string {
	switch d {
	case point_a:
		return "POINT_A"
	case point_b:
		return "POINT_B"
	case point_c:
		return "POINT_C"
	case point_d:
		return "POINT_D"
	default:
		return "NOTHING"
	}
}

func (d Direction) getNameV2() string {
	return []string{"POINT_A", "POINT_B", "POINT_C", "POINT_D"}[d]
}

func main() {
	var x Direction = 4
	fmt.Println(x.getName())
	fmt.Println(point_a.getName())
	fmt.Println(point_c.getName())
}
