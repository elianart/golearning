package main

import (
	"fmt"
)

func sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 1000; i++ {
		z = calc(z, x)

	}
	return z
}

func calc(z float64, x float64) float64 {
	z -= (z*z - x) / (2 * z)
	return z
}

func main() {
	fmt.Println(sqrt(2))
}
