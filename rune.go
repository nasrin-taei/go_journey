package main

import "fmt"

func main() {
	var c rune

	c = 'A'
	s := "character"

	out := fmt.Sprintf("%c %s\n", c, s)
	fmt.Print(out)
}
