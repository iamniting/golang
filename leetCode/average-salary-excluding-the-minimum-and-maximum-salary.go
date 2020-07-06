package main

import "math"

func average(salary []int) float64 {

	var sum int
	min, max := math.MaxInt32, math.MinInt32

	for _, s := range salary {
		min = int(math.Min(float64(min), float64(s)))
		max = int(math.Max(float64(max), float64(s)))

		sum += s
	}

	return float64(sum-min-max) / float64(len(salary)-2)
}
