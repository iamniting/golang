package main

import "sort"

func canMakeArithmeticProgression(arr []int) bool {

	sort.Ints(arr)
	diff := arr[1] - arr[0]

	for i := range arr[1:] {
		if arr[i+1]-arr[i] != diff {
			return false
		}
	}

	return true
}
