package main

import "math"

func largestAltitude(gain []int) int {

	var sum, max int

	for _, num := range gain {
		sum += num
		max = int(math.Max(float64(sum), float64(max)))
	}

	return max
}
