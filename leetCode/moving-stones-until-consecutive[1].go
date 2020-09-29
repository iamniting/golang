package main

import "sort"

func numMovesStones(a int, b int, c int) []int {

	stones := []int{a, b, c}
	sort.Ints(stones)
	a, b, c = stones[0], stones[1], stones[2]

	if c-a == 2 {
		return []int{0, 0}
	}

	if c-b > 2 && b-a > 2 {
		return []int{2, c - a - 2}
	}

	return []int{1, c - a - 2}
}
