package main

import (
	"fmt"
	"unicode"
)

func countDifferentIntegers(word string) int {
	result := 0
	for i := 0; i < len(word); i++ {
		if unicode.IsDigit(rune(word[i])) {
			count := 1
			for j := i + 1; j < len(word); j++ {
				if unicode.IsDigit(rune(word[j])) {
					count++
				} else {
					break
				}
			}
			i += count
			result++
		}
	}
	return result
}

func main() {
	word := "a123bc34d8ef34"
	result := countDifferentIntegers(word)
	fmt.Println(result)
}
