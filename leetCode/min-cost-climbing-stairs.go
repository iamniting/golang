package main

import "math"

func minCostClimbingStairs(cost []int) int {

	var f1, f2 int

	for i := len(cost) - 1; i >= 0; i-- {
		f1, f2 = cost[i]+int(math.Min(float64(f1), float64(f2))), f1
	}

	return int(math.Min(float64(f1), float64(f2)))
}
