package main

import "math"

func maxScore(s string) int {
	var ones, one, zero, res int

	for _, c := range s {
		if c == '1' {
			ones++
		}
	}

	for _, c := range s[:len(s)-1] {
		if c == '1' {
			one++
		} else {
			zero++
		}

		res = int(math.Max(float64(res), float64(zero+(ones-one))))
	}

	return res
}
