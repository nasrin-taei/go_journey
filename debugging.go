package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		for i := 0; ; i++ {
			time.Sleep(1 * time.Second)
			fmt.Println(i)
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("Finished")
}
