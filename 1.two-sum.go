package main

func twoSum(nums []int, target int) []int {

	// my approach
	result := []int{}
	for left := 0; left <= len(nums); left++ {
		for right := left + 1; right < len(nums); right++ {
			if nums[left]+nums[right] == target {
				result = append(result, left, right)
				return result
			}
		}
	}

	// Better approach with O(n) complexity

	// diffToIndex := make(map[int]int, 0)
	//
	// outputIndices := []int{}
	//
	// for i, num := range nums {
	// 	diffToIndex[target-num] = i
	// }
	//
	// for firstIndex, num := range nums {
	// 	secondIndex, exists := diffToIndex[num]
	//
	// 	if exists && firstIndex != secondIndex {
	// 		outputIndices = append(outputIndices, firstIndex)
	// 		outputIndices = append(outputIndices, secondIndex)
	// 		return outputIndices
	// 	}
	// }

	return nil
}
