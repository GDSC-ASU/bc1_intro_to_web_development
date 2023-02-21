package main

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(n int) (float64, error) {
	if n < 0 {
		return 0.0, errors.New("can't take square root of a negative number")
	}

	return math.Sqrt(float64(n)), nil
}

func main() {
	value, err := sqrt(-12)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(value)
}
