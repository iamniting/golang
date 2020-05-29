package main

import "math"

func minTimeToVisitAllPoints(points [][]int) int {

	getDiff := func(p1 []int, p2 []int) int {

		xDiff := math.Abs(float64(p1[0] - p2[0]))
		yDiff := math.Abs(float64(p1[1] - p2[1]))

		return int(math.Max(xDiff, yDiff))
	}

	diff := 0
	for i := 1; i < len(points); i++ {
		diff += getDiff(points[i-1], points[i])
	}
	return diff
}
