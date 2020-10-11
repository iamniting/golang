package main

func maxDepth(s string) int {

	var stack []rune
	var res int

	for _, c := range s {
		if c == '(' {
			stack = append(stack, c)
			if len(stack) > res {
				res = len(stack)
			}
		} else if c == ')' {
			stack = stack[:len(stack)-1]
		}
	}

	return res
}
