package main

import (
	"math"
	"sort"
)

func leastBricks(wall [][]int) int {

	res := len(wall)
	m := make([]map[int]bool, len(wall))
	common := map[int]bool{}

	// make a map of wall
	for i, row := range wall {
		m[i] = map[int]bool{}
		var sum int

		for _, c := range row[:len(row)-1] {
			sum += c
			m[i][sum] = true
			common[sum] = true
		}
	}

	// iterate through all common sums and find them in map of wall
	for key := range common {
		var cuts int
		for i := 0; i < len(wall); i++ {
			if !m[i][key] {
				cuts++
			}
		}
		res = int(math.Min(float64(res), float64(cuts)))
	}

	return res
}
