package main

import (
	"math"
	"sort"
)

func largestSumAfterKNegations(A []int, K int) int {

	sort.Ints(A)

	// run only if -ve nums
	for i := 0; i < len(A) && K > 0 && A[i] < 0; i++ {
		A[i] = -A[i]
		K--
	}

	// if k is not 0 find minimum and change it to -ve
	if K > 0 && K&1 == 1 {

		min := math.MaxInt32
		minIndex := -1
		for i, n := range A {
			if n < min {
				min = n
				minIndex = i
			}
		}

		A[minIndex] = -A[minIndex]
	}

	var sum int
	for _, n := range A {
		sum += n
	}

	return sum
}
