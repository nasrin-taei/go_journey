package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int, 6)

	go func() {
		for {
			time.Sleep(5 * time.Second)
			x, ok := <-channel
			if !ok {
				break
			}
			time.Sleep(10 * time.Millisecond)
			fmt.Println("Received value is ", x)
		}
	}()

	for i := 0; i < 10; i++ {
		channel <- 10 + i
		fmt.Println("Sent value is ", 10+i)
	}
	close(channel)

	time.Sleep(33 * time.Second)
	fmt.Println("finished!!!!!!!!!!")
}
