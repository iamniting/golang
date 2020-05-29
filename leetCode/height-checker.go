package main

import "sort"

func heightChecker(heights []int) int {

	sorted := make([]int, len(heights))
	copy(sorted[:], heights)
	sort.Ints(sorted)

	res := 0
	for i := range heights {
		if sorted[i] != heights[i] {
			res++
		}
	}

	return res
}
