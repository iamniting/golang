package main

import "sort"

func sortByBits(arr []int) []int {

	sort.Slice(arr, func(i, j int) bool {
		x, y := getBitsCount(arr[i]), getBitsCount(arr[j])
		if x == y {
			return arr[i] < arr[j]
		}
		return x < y
	})
	return arr
}

func getBitsCount(n int) int {
	res := 0

	for n > 0 {
		res += n & 1
		n >>= 1
	}
	return res
}
