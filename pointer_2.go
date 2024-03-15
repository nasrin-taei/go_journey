package main

import "fmt"

type myStruct struct {
	a int
	s string
}

func main() {
	newStruct := myStruct{
		a: 0,
		s: "",
	}
	fmt.Println(newStruct)
	updateStruct(&newStruct)
	fmt.Println(newStruct)

	mySlice := []int{1, 3}
	fmt.Println(mySlice)
	updateSlice(&mySlice)
	fmt.Println(mySlice)
}

func updateStruct(obj *myStruct) {
	obj = &myStruct{
		a: 0,
		s: "",
	}
	obj.s = "updated"
	obj.a = 100
}

func updateSlice(slice *[]int) {
	tmp := *slice
	tmp[0] = 10
	tmp[1] = 2
}
