package main

import "math"

func findLengthOfLCIS(nums []int) int {

	var res, prevIndex int
	prev := math.MinInt32

	for i, n := range append(nums, math.MinInt32) {
		if n <= prev {
			res = int(math.Max(float64(res), float64(i-prevIndex)))
			prevIndex = i
		}
		prev = n
	}
	return res
}
