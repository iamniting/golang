package main

import "sort"

func trimMean(arr []int) float64 {
	sort.Ints(arr)

	count := len(arr) / 20

	var sum int
	for i := count; i < len(arr)-count; i++ {
		sum += arr[i]
	}

	return float64(sum) / float64(len(arr)-count-count)
}
