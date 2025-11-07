package main

// Constraints:
//
// nums1.length == m + n
// nums2.length == n
// 0 <= m, n <= 200
// 1 <= m + n <= 200
// -109 <= nums1[i], nums2[j] <= 109
//
// Approach:
// work backward from nums1 since nums1 has length of m + n
// Example: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
//
// Compare nums1[2]=3 vs nums2[2]=6 → Put 6 at nums1[5]
// Compare nums1[2]=3 vs nums2[1]=5 → Put 5 at nums1[4]
// Compare nums1[2]=3 vs nums2[0]=2 → Put 3 at nums1[3]
// Compare nums1[1]=2 vs nums2[0]=2 → Put 2 at nums1[2]
// Compare nums1[0]=1 vs nums2[0]=2 → Put 2 at nums1[1]
// Result: [1,2,2,3,5,6]

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	k := m + n - 1

	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

}
