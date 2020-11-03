package main

import (
	"math"
	"sort"
)

func maxWidthOfVerticalArea(points [][]int) int {

	sort.Slice(points, func(i, j int) bool { return points[i][0] < points[j][0] })

	var res float64

	for i := 1; i < len(points); i++ {
		res = math.Max(res, float64(points[i][0]-points[i-1][0]))
	}

	return int(res)
}
