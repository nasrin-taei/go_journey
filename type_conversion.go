package main

import "fmt"

type myInt int32

func main() {

	b := make([]byte, 10)
	s := string(b)

	fmt.Println(s)

	b = []byte(s)

	i := myInt(10)
	i2 := int32(i)
	i = myInt(i2)

	fmt.Printf("%d = %d", i, i2)
}
