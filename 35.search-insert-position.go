package main

// Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
//
// You must write an algorithm with O(log n) runtime complexity.
//
//
// Aprroach:
//
// use Binary Search algorithm
//

func searchInsert(nums []int, target int) int {
	i, j := 0, len(nums)-1

	for i <= j {
		mid := i + (j-i)/2

		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			i = mid + 1
			continue
		} else {
			j = mid - 1
			continue
		}
	}
	return i
}
