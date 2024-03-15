package main

import (
	"fmt"
	"time"
)

func main() {
	chanel := make(chan int)

	go func() {
		defer close(chanel)
		for i := 0; i < 4; i++ {
			chanel <- 2 * i
			fmt.Printf("sent value is %v\n", 2*i)
		}
	}()

	go func() {
		for {
			x, ok := <-chanel
			if !ok {
				break
			}
			time.Sleep(1 * time.Second)
			fmt.Println("received value is", x)
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("Bye")

}
