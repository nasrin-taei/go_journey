package main

import "fmt"

func main() {
	defer saySomething("hello from defer....")

	print("Hiiii")

	defer func(s string) {
		fmt.Println(s)
	}("Bye defer")

	defer saySomething("bye byeeee")

	defer func(s string) {
		fmt.Println(s)
	}("Bye3")
}

func print(s string) {
	fmt.Println(s)
}

func saySomething(s string) {
	fmt.Println(s)
}
