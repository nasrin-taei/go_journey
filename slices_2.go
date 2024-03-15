package main

import (
	"fmt"
)

func main() {

	myArray := [5]int{2, 8, 6, 8, 110}
	fmt.Println(len(myArray), cap(myArray))

	mySlice := myArray[2:3]
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

	fmt.Println(myArray)
	fmt.Println(mySlice)

	myArray[2] = 10
	fmt.Println(myArray)
	fmt.Println(mySlice)

	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 2)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

	mySlice = append(mySlice, 2)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

	fmt.Println(myArray)
	fmt.Println(mySlice)

	mySlice[0] = 11
	fmt.Println(myArray)
	fmt.Println(mySlice)

	//mySlice = append(mySlice, 10)
	//fmt.Println(mySlice)
	//fmt.Println(len(mySlice))
	//fmt.Println(cap(mySlice))
	//
	//mySlice = append(mySlice, 20)
	//fmt.Println(mySlice)
	//fmt.Println(len(mySlice))
	//fmt.Println(cap(mySlice))
	//
	//mySlice = append(mySlice, 70)
	//fmt.Println(mySlice)
	//fmt.Println(len(mySlice))
	//fmt.Println(cap(mySlice))

}
