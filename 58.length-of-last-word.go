package main

import "strings"

func lengthOfLastWord(s string) int {
	words := strings.Split(strings.TrimSpace(s), " ")

	return len(words[len(words)-1])
}
