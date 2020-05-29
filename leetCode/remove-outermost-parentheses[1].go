package main

func removeOuterParentheses(S string) string {
	start, index := 0, 0
	res := ""

	for i, c := range S {
		if c == '(' {
			index++
		} else {
			index--
		}

		if index == 0 {
			res += S[start+1 : i]
			start = i + 1
		}
	}
	return res
}
