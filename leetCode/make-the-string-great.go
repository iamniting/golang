package main

import "math"

func makeGood(s string) string {

	var stack []byte

	for i := range s {

		if len(stack) > 0 && int(math.Abs(float64(stack[len(stack)-1])-float64(s[i]))) == 32 {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return string(stack)
}
