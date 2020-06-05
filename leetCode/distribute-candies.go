package main

import "math"

func distributeCandies(candies []int) int {
	set := map[int]interface{}{}

	for _, c := range candies {
		set[c] = nil
	}

	kind, sis := len(set), len(candies)>>1

	return int(math.Min(float64(kind), float64(sis)))
}
