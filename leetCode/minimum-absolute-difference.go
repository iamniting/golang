package main

import (
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {

	min := math.MaxInt32
	res := [][]int{}

	sort.Ints(arr)

	// get minimum  diff
	for i := 1; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]
		if diff < min {
			min = diff
		}
		if diff == 1 {
			break
		}
	}

	// create pairs with minimum diff
	for i := 1; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]
		if diff == min {
			res = append(res, []int{arr[i-1], arr[i]})
		}
	}

	return res
}
