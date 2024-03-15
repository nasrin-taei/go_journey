package main

import "fmt"

func surround(msg string, left, right rune) string {
	return fmt.Sprintf("%c%v%c", left, msg, right)

}

func main() {

	fmt.Printf("Hello,world! \n")
	fmt.Printf("%v,%v!\n", "Hello", "world")
	fmt.Printf("This is a \"Quote\" \n")

}
