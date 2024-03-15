package main

import "fmt"

func main() {

	myArray := [4]int{1, 2, 3, 4}
	fmt.Println(myArray)

	mySlice := make([]int, 0, 4)
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

	mySlice2 := []int{1, 2, 3}
	fmt.Println(mySlice2)
	fmt.Println(len(mySlice2))
	fmt.Println(cap(mySlice2))

	mySlice2[2] = 4
	fmt.Println(mySlice2)

	mySlice2 = append(mySlice2, 20)
	fmt.Println(mySlice2)
	fmt.Println(len(mySlice2))
	fmt.Println(cap(mySlice2))

	mySlice[3] = 21

	var mySlice3 []int
	fmt.Println(mySlice3)
	fmt.Println(len(mySlice3))
	fmt.Println(cap(mySlice3))

	mySlice3 = append(mySlice3, 10)
	fmt.Println(mySlice3)
	fmt.Println(len(mySlice3))
	fmt.Println(cap(mySlice3))

	mySlice3[0] = 2
	fmt.Println(mySlice3)
	fmt.Println(len(mySlice3))
	fmt.Println(cap(mySlice3))

}
