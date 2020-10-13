package main

import "sort"

func maxCoins(piles []int) int {

	sort.Ints(piles)
	var sum int

	for i := len(piles) - 2; i >= len(piles)/3; i -= 2 {
		sum += piles[i]
	}

	return sum
}
