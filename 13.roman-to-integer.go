package main

import (
	"fmt"
)

func romanToInteger(s string) int {
	roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	runes := []rune(s)
	n := 0
	lastNumber := 0

	for i := len(runes) - 1; i >= 0; i-- {
		current := roman[string(runes[i])]

		if lastNumber > current {
			n -= current
		} else {
			n += current
		}

		lastNumber = current

	}

	return n
}

func main() {
	nomer := romanToInteger("IV")

	fmt.Println("nomer: ", nomer)
}
