package main

import "math"

func maxCount(m int, n int, ops [][]int) int {

	for _, o := range ops {
		m = int(math.Min(float64(m), float64(o[0])))
		n = int(math.Min(float64(n), float64(o[1])))
	}

	return m * n
}
