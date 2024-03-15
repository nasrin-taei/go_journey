package main

import (
	"errors"
	"fmt"
)

type SomeError struct {
	code int
}

func (s SomeError) Error() string {
	return "some error happened"
}

func someFunc() (int, error) {
	return 0, SomeError{code: 1}
}

func main() {
	i, err := someFunc()
	if err != nil {
		var someErr SomeError
		if errors.Is(err, SomeError{code: 2}) {
			fmt.Println(err)
			return
		} else if errors.As(err, &someErr) {
			fmt.Println("Error is not critical but is important which has code", someErr.code)
		} else {
			fmt.Println("Do not care about the error")
		}
	}
	fmt.Println("i is", i)
}
