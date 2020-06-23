package main

import "math"

func findShortestSubArray(nums []int) int {

	freq, left, right := map[int]int{}, map[int]int{}, map[int]int{}
	degree := 0

	for i, n := range nums {
		if _, ok := left[n]; !ok {
			left[n] = i
		}
		right[n] = i
		freq[n]++
		degree = int(math.Max(float64(degree), float64(freq[n])))
	}

	minLen := math.MaxInt32
	for k, v := range freq {
		if v == degree {
			minLen = int(math.Min(float64(minLen), float64(right[k]-left[k]+1)))
		}
	}
	return minLen
}
