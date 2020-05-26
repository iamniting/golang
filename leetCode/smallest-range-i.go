package main

import "math"

func smallestRangeI(A []int, K int) int {
	min, max := math.MaxInt32, 0

	for _, n := range A {
		min = int(math.Min(float64(min), float64(n)))
		max = int(math.Max(float64(max), float64(n)))
	}

	min = min + K
	max = max - K

	if min > max {
		return 0
	}
	return max - min
}
