package main

import "math"

func findShortestSubArray(nums []int) int {

	m := map[int][]int{}
	degree := 0

	for i, n := range nums {
		m[n] = append(m[n], i)
		degree = int(math.Max(float64(degree), float64(len(m[n]))))
	}

	minLen := math.MaxInt32
	for _, v := range m {
		if len(v) == degree {
			fIndex, lIndex := v[0], v[len(v)-1]
			minLen = int(math.Min(float64(minLen), float64(lIndex-fIndex+1)))
		}
	}
	return minLen
}
