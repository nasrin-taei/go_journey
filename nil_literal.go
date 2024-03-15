package main

import "fmt"

type myInterface interface {
	myPrint()
}

func main() {

	var myMap map[string]int //must make map.
	if myMap == nil {
		fmt.Printf("%#v\n", myMap)
	}
	fmt.Println(myMap[""]) //Prints 0 cause map is nil

	var mySlice []int //must make slice.
	if mySlice == nil {
		fmt.Printf("%#v\n", mySlice)
	}
	//fmt.Println(mySlice[0]) cause runtime error

	var myInterfaceVar myInterface //must be initialized with its implementation.
	if myInterfaceVar == nil {
		fmt.Printf("%#v\n", myInterfaceVar)
	}
	//myInterfaceVar.myPrint()  cause runtime error
}
