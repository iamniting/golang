package main

import "math"

func findLHS(nums []int) int {

	var res int
	m := map[int]int{}

	for _, n := range nums {
		m[n]++
	}

	for _, n := range nums {
		// n, n-1 combination
		if val, ok := m[n-1]; ok {
			res = int(math.Max(float64(res), float64(m[n]+val)))
		}
		// n, n+1 combination
		if val, ok := m[n+1]; ok {
			res = int(math.Max(float64(res), float64(m[n]+val)))
		}
	}

	return res
}
