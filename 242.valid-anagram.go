package main

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := map[rune]int{}

	for _, v := range s {
		sMap[v]++
	}

	for _, l := range t {
		if sMap[l] == 0 {
			return false
		}
		sMap[l]--
	}

	return true
}
