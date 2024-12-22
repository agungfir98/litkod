package main

func validParentheses(s string) bool {
	pts := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}

	if len(s)%2 != 0 {
		return false
	}

	runes := []rune(s)
	stack := []string{}

	for _, rune := range runes {
		if len(stack) > 0 {
			lastStack := stack[len(stack)-1]
			pairs := pts[lastStack]

			if pairs == string(rune) {
				stack = stack[:len(stack)-1]
				continue
			}
		}
		stack = append(stack, string(rune))
	}

	if len(stack) == 0 {
		return true
	}

	return false
}
