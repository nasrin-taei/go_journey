package main

import "fmt"

func main() {

	var mySlice []int
	mySlice = []int{2, 4, 5}
	fmt.Println(mySlice)
	fmt.Println(mySlice[2])
	mySlice[0] = 10

	mySlice = append(mySlice, 6)
	fmt.Println(mySlice)

	//var myMap map[string]int
	myMap := map[string]int{"Nasrin": 34, "Behrad": 32, "Khar": 40}
	fmt.Println(myMap)
	fmt.Println(myMap["Behrad"])

	//Update
	myMap["Behrad"] = 33
	fmt.Println(myMap["Behrad"])

	//Add
	myMap["Naghi"] = 50
	fmt.Println(myMap)

	//Add with duplicate VALUE
	myMap["Zahra"] = 40
	fmt.Println(myMap)

	for k, v := range myMap {
		fmt.Println(k, v)
	}

	delete(myMap, "Khar")
	fmt.Println(myMap)

	//Check existence
	age := myMap["Zahra"]
	fmt.Println("Zahra :", age)

	age = myMap["Hossein"]
	fmt.Println("Hossein :", age)

	age, existed := myMap["Hossein"]
	if existed {
		fmt.Println("(Check existence) Hossein :", age)
	}

	age, existed = myMap["Zahra"]
	if existed {
		fmt.Println("(Check existence) Zahra :", age)
	}
}
