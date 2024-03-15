package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan int)

	go func() {
		//defer close(channel)
		for i := 0; i < 5; i++ {
			time.Sleep(1 * time.Second)
			channel <- 10 + i
			fmt.Println("Sent value is ", 10+i)
		}
		close(channel)
	}()

	for {
		x, ok := <-channel
		if !ok {
			break
		}
		fmt.Println("Received value is ", x)
		time.Sleep(5 * time.Second)
	}
}
