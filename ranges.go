package main

import "fmt"

func main() {

	mySlice := make([]string, 3, 5)
	mySlice = []string{"Behrad", "Nasrin", "Hava"}

	//Iterating over an array or slice using normal for and indexes
	for i := 0; i < len(mySlice); i++ {
		fmt.Println(i, mySlice[i])
	}

	fmt.Println("------------------------------------")

	//Using range keyword in two ways
	//1
	for i := range mySlice {
		fmt.Println(i, mySlice[i])
	}

	//2
	fmt.Println("------------------------------------")
	for i, e := range mySlice {
		fmt.Println(i, e)
	}
}
