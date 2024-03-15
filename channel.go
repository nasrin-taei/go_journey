package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(1 * time.Second)
			channel <- 10 + i
			fmt.Println("Sent value is ", 10+i)
		}
	}()

	for i := 0; i < 5; i++ {
		x, ok := <-channel
		if !ok {
			break
		}
		time.Sleep(1 * time.Second)
		fmt.Println("Received value is ", x)
	}

}
