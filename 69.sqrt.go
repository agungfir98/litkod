package main

func sqrt(x int) int {
	l, r := 0, x

	var result int

	for l <= r {
		mid := l + (r-l)/2
		pow := mid * mid

		if x == pow {
			return mid
		} else if x > pow {
			result = mid
			l = mid + 1
		} else {
			r = mid - 1
		}

	}

	return result
}
