package main

import (
	"fmt"
	"time"
)

func count(amount int) {
	for i := 1; i <= amount; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("count", i)
	}
}

func main() {
	myAmount := 5

	go count(5)

	go func(amount int) {
		for i := 1; i <= myAmount; i++ {
			time.Sleep(1000 * time.Millisecond)
			fmt.Println("literal", i)
		}
	}(5)

	fmt.Println("wait for goroutine")
	time.Sleep(10000 * time.Millisecond)
	fmt.Println("end program")
}
