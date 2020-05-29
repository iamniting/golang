// https://www.youtube.com/watch?v=r0-zx5ejdq0

package main

import "math"

func getMaxDiff(slice []int, l int) int {
	max := 0

	slice = append(slice, l+1)
	slice = append([]int{0}, slice...)

	for i := 0; i < len(slice)-1; i++ {
		max = int(math.Max(float64(max), float64(slice[i+1]-slice[i]-1)))
	}
	return max
}

func longestValidParentheses(s string) int {
	var stack = []int{}

	for i, c := range s {
		l := len(stack)

		if c == '(' {
			stack = append(stack, i+1)
		} else if c == ')' && len(stack) > 0 && s[stack[l-1]-1] == '(' {
			stack = stack[:len(stack)-1]
		} else if c == ')' && len(stack) == 0 || s[stack[l-1]-1] == ')' {
			stack = append(stack, i+1)
		}
	}

	if len(stack) == 0 {
		return len(s)
	}

	return getMaxDiff(stack, len(s))
}
