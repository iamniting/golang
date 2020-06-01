package main

import "math"

func minCostToMoveChips(chips []int) int {
	odd, even := 0, 0

	for _, n := range chips {
		if n&1 == 1 {
			odd++
		} else {
			even++
		}
	}

	return int(math.Min(float64(odd), float64(even)))
}
