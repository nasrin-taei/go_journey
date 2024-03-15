package main

import (
	"fmt"
	"math/rand"
)

type funcAlias func(s string) string //CustomType or TypeAlias for a function

func printAsYouWant(makeMyName func(name string) string, myString string) {
	fmt.Println(makeMyName(myString))
}

func makeSomeFunction() func(int) int {
	return func(i int) int {
		return i + 1
	}
}

func main() {

	//Below anonymous function passed to myFunc as literal
	var myFunc func(string) string
	msg := "It's my name :"
	myFunc = func(name string) string {
		return fmt.Sprint(msg, name)
	}
	s := myFunc("behrad")
	fmt.Println(s)

	//Used as another function parameter
	printAsYouWant(func(name string) string {
		return name
	}, "Asghar")
	printAsYouWant(myFunc, "Nasrin")

	//Calling anonymous function / literal
	x := func(a, b int) int {
		return a + b
	}(2, 3)
	fmt.Println(2, "+", 3, "=", x)

	x2 := func() int {
		return rand.Intn(14)
	}()
	fmt.Println(x2)

	//printAsYouWant function passed as literal
	myFunc2 := printAsYouWant

	myFunc2(myFunc, "Noor Norouzieh")

	//Returning function
	myFunc3 := makeSomeFunction() //farakhani
	fmt.Println(myFunc3(8))

	//Functions as Types
	fn4 := funcAlias(func(s string) string {
		return s
	})
	fmt.Println(fn4("Giti"))

	var fn5 funcAlias
	fn5 = func(s string) string {
		return s
	}
	fmt.Println(fn5("Whatever"))
}
