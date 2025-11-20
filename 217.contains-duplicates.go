package main

func containsDuplicate(nums []int) bool {
	numMap := map[int]bool{}

	for _, v := range nums {
		if numMap[v] {
			return true
		}
		numMap[v] = true
	}

	return false
}
