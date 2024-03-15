package main

import "fmt"

func main() {
	mySlice := make([]string, 0, 4)
	printSliceElementsWithAddress(mySlice)

	addToSlice(mySlice, "OK")
	printSliceElementsWithAddress(mySlice)

	addToSlice(mySlice, "OK")
	addToSlice(mySlice, "OK")
	addToSlice(mySlice, "OK")
	printSliceElementsWithAddress(mySlice)

	addToSlice(mySlice, "OK")
	printSliceElementsWithAddress(mySlice)
}

func printSliceElementsWithAddress(slice []string) {
	for i := range slice {
		fmt.Printf("slice element %d = \"%s\" at %p\n", i, slice[i], &slice[i])
	}
	fmt.Println("-------------------------------------")
}

func addToSlice(slice []string, element string) {
	slice = append(slice, element)
	fmt.Printf("len(%d) and cap(%d) and %s is added to slice\n", len(slice), cap(slice), element)
}
