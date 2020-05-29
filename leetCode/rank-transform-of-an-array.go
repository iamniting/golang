package main

import "sort"

func arrayRankTransform(arr []int) []int {
	// copy the arr in res, and use res as tmp array to sort elements
	res := append([]int{}, arr...)
	m := make(map[int]int)
	sort.Ints(res)

	// put rank in the map
	rank := 0
	for _, n := range res {
		if _, ok := m[n]; !ok {
			rank++
			m[n] = rank
		}
	}

	// copy the rank in res from the map
	for i, n := range arr {
		res[i] = m[n]
	}
	return res
}
