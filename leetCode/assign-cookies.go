package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	var child, cookies int
	for cookies = 0; cookies < len(s) && child < len(g); cookies++ {
		if s[cookies] >= g[child] {
			child++
		}
	}
	return child
}
