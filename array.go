package main

import "fmt"

func printArray(title string, slice [5]string) {
	fmt.Println(title)
	for i := 0; i < len(slice); i++ {
		element := slice[i]
		fmt.Println(element)
	}
}

func main() {
	route := [5]string{"grocery", "department store", "salon"}
	printArray("route1 is", route)
	route[4] = "home"
	printArray("route2 is", route)
	fmt.Println(route[0], "visited")
	fmt.Println(route[1], "visited")
	printArray("preparing route is", route)
}
