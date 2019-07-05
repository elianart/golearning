package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	array := strings.Split(s, " ")
	wordArray := make(map[string]int)
	for _, s := range array {
		val := wordArray[s]
		if val != 0 {
			val++
			wordArray[s] = val
		} else {
			wordArray[s] = 1
		}
	}

	return wordArray
}

func main() {
	wc.Test(WordCount)
}
