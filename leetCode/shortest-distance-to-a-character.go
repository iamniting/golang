package main

import "math"

func shortestToChar(S string, C byte) []int {

	res := make([]int, len(S))

	for i := range S {
		min := math.MaxInt32

		// check before the current element
		for j := i; j >= 0; j-- {
			if S[j] == C && i-j < min {
				min = i - j
			}
		}

		// check after the current element
		for j := i + 1; j < len(S); j++ {
			if S[j] == C && j-i < min {
				min = j - i
			}
		}

		res[i] = min
	}
	return res
}
