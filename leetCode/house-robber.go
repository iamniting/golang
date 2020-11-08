package main

import "math"

func rob(nums []int) int {

	var prevprev, prev int

	for _, n := range nums {
		prevprev, prev = prev, int(math.Max(float64(prevprev+n), float64(prev)))
	}

	return prev
}
