package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(2 * time.Second)
			fmt.Println("Hello from go routine", i)
		}()
	}

	wg.Wait()
}
