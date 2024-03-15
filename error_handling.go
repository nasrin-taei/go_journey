package main

import (
	"errors"
	"fmt"
)

type Division struct {
	a, b float64
}

func (d Division) divide() (float64, error) { //Raising error!!!
	if d.a == 0 && d.b == 0 {
		return 0, errors.New("two numbers are zero")
	} else if d.b == 0 {
		return 0, DivisionError{num: d.a, callerName: "divide"}
	} else {
		return d.a / d.b, nil
	}
}

func main() {
	for a := float64(0); a < 4; a++ {
		for b := float64(0); b < 4; b++ {
			result, err := Division{a: a, b: b}.divide()
			if err != nil { //Handling the error
				fmt.Println(err)
				continue
			}
			fmt.Printf("%.2f / %.2f = %.2f\n", a, b, result)
		}
	}
	fmt.Println("Loop ended.")

}

type DivisionError struct {
	num        float64
	callerName string
}

func (t DivisionError) Error() string {
	return fmt.Sprintf("Number %0.2f is not dividable by zero in function %s", t.num, t.callerName)
}
