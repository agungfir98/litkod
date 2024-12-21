package main

func isPalindrome(x int) bool {
	if x <= 0 {
		return false
	}

	original := x
	reversed := 0

	for 0 < x {
		digit := x % 10
		reversed = reversed*10 + digit
		x /= 10
	}

	if original == reversed {
		return true
	}

	return false
}
