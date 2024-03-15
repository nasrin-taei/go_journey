package main

import "fmt"

func main() {
	s1 := []int{8, 5, 6}
	s2 := []int{2, 4, 7, 9, 3, 1}
	copy(s2, s1)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(&s1[0])
}
