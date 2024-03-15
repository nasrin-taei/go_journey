package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	mx := sync.Mutex{}

	sum := 0
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			fmt.Println("Hello from GO ROUTINE ", i)
			mx.Lock()
			sum = sum + 1
			mx.Unlock()
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("sum is", sum)
}
