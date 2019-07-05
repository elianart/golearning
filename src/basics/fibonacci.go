package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	previous := make([]int, 10)
	return func(x int) int {
		result := 0
		len := len(previous)
		if x == 0 || x == 1 {
			result = x
		}
		if x == 2 {
			result = 1
		} else {
			result = previous[len-1] + previous[len-2]
		}

		previous = append(previous, result)
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
