package main

import "sort"

func maxCoins(piles []int) int {

	sort.Ints(piles)
	var sum int
	i, j := 0, len(piles)-2

	// i for Bob, j for myself, j+1 for Alice
	for i < j {
		sum += piles[j]
		i++
		j -= 2
	}

	return sum
}
