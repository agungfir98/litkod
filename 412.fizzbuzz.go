package main

import "strconv"

func fizzbuzz(n int) []string {
	result := make([]string, 0, n)

	for i := range n {
		i++
		if i%3 == 0 && i%5 == 0 {
			result = append(result, "fizzbuzz")
			continue
		}
		if i%3 == 0 {
			result = append(result, "fizz")
			continue
		}
		if i%5 == 0 {
			result = append(result, "buzz")
			continue
		}

		result = append(result, strconv.Itoa(i))
	}

	return result

}
