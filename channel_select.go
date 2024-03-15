package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan int)
	c2 := make(chan int, 2)

	go func() {
		//defer close(c1)
		i := 0
		for {
			time.Sleep(500 * time.Millisecond)
			c1 <- i
			fmt.Println("sent value to c1 is", i)
			i++
		}
		close(c1)
	}()

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			x, ok := <-c2
			if !ok {
				break
			}
			fmt.Println("received value from c2 is", x)
		}
	}()

	for i := 0; i < 8; i++ {
		select {
		case x := <-c1:
			time.Sleep(1 * time.Second)
			fmt.Println("select received value from c1 is", x)
		case c2 <- i:
			time.Sleep(1 * time.Second)
			fmt.Println("select sent value to c2 is", i)
		case <-time.After(100 * time.Millisecond):
			fmt.Println("yohooooooooooooooooooo time out")
			return

		}

	}
	fmt.Println("finish")
}
