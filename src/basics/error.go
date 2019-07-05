package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt struct {
	value float64
	text  string
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("%s: %v", e.text, e.value)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		error := &ErrNegativeSqrt{
			x,
			"cannot Sqrt negative number",
		}
		return 0, error
	}
	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
