package main

// missing number = n(n+1)/2
func missingNumber(nums []int) int {
	n := len(nums)
	expectedNumber := (n * (n + 1)) / 2
	total := 0

	for _, v := range nums {
		total += v
	}

	return expectedNumber - total
}
