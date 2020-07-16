package main

import "math"

func minMoves(nums []int) int {
	sum, min := 0, math.MaxInt32

	for _, num := range nums {
		sum += num
		min = int(math.Min(float64(min), float64(num)))
	}

	return sum - min*len(nums)
}
