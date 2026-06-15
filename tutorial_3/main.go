package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("divideFloat(46, 3):", divideFloat(46, 3))

	result, remainder, err := divideInt(46, 3)
	fmt.Printf("divideInt(46, 3): result=%v, remainder=%v, error=%v\n", result, remainder, err)

	result2, remainder2, err2 := divideInt(67, 0)
	fmt.Printf("divideInt(67, 0): result=%v, remainder=%v, error=%v\n", result2, remainder2, err2)

	if err2 != nil {
		fmt.Println("the second division operation resulted in an error:", err2)
	} else {
		switch remainder {
		case 0:
			fmt.Println("the second division was exact")
		case 1, 2:
			fmt.Println("the second division was close")
		default:
			fmt.Println("the second division was not close")
		}
	}
}

func divideFloat(numerator float64, denominator float64) float64 {
	return numerator / denominator
}

func divideInt(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}

	result := numerator / denominator
	remainder := numerator % denominator
	return result, remainder, nil
}
