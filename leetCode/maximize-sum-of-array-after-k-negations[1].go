package main

import (
	"math"
	"sort"
)

func largestSumAfterKNegations(A []int, K int) int {

	sort.Ints(A)
	min := math.MaxInt32
	var sum int

	for i := 0; i < len(A) && K > 0 && A[i] < 0; i++ {
		A[i] = -A[i]
		K--
	}

	for _, n := range A {
		sum += n
		min = int(math.Min(float64(min), float64(n)))
	}

	// minus  (2 * min) from sum is K is odd
	return sum - ((K & 1) * min * 2)
}
