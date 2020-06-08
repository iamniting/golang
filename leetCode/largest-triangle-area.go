package main

import "math"

func largestTriangleArea(points [][]int) float64 {
	cal := func(x1, y1, x2, y2, x3, y3 int) float64 {
		return math.Abs(float64(x1*y2+x2*y3+x3*y1-y1*x2-y2*x3-y3*x1)) / 2.0
	}

	var max float64
	for i := 0; i < len(points)-2; i++ {
		for j := i + 1; j < len(points)-1; j++ {
			for k := j + 1; k < len(points); k++ {
				max = math.Max(max, cal(
					points[i][0], points[i][1],
					points[j][0], points[j][1],
					points[k][0], points[k][1]))
			}
		}
	}
	return max
}
