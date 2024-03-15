package main

import "fmt"

func main() {
	//
	//user := "nasrin"
	//age := 35
	//
	//if user == "nasrin" && age >= 20 {
	//	fmt.Println("Show something +20 only to Nasrin")
	//} else if age >= 20 {
	//	fmt.Println("Show some public +20 content ")
	//} else {
	//	fmt.Println("Got nothing for you. please get back when you're at least 20.")
	//}

	if age, user := getAgeAndName(); user == "nasrin" && age >= 20 {
		fmt.Println("Show something +20 only to Nasrin")
	} else if age >= 20 {
		fmt.Println("Show some public +20 content ")
	} else {
		fmt.Println("Got nothing for you. please get back when you're at least 20.")
	}

	fmt.Println("End of the program")
}

func getAgeAndName() (int, string) {
	return 12, "nain"
}
