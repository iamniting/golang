package main

import "math"

func findShortestSubArray(nums []int) int {

	m := map[int]int{}
	degree := 0

	for _, n := range nums {
		m[n]++
		degree = int(math.Max(float64(degree), float64(m[n])))
	}

	minLen := math.MaxInt32
	for k, v := range m {
		if v != degree {
			continue
		}

		var fIndex, lIndex int
		for i, n := range nums {
			if n == k {
				fIndex = i
				break
			}
		}
		for i := len(nums) - 1; i >= 0; i-- {
			if nums[i] == k {
				lIndex = i
				break
			}
		}
		minLen = int(math.Min(float64(minLen), float64(lIndex-fIndex+1)))
	}
	return minLen
}
