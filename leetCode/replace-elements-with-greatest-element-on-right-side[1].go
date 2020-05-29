package main

import "math"

func replaceElements(arr []int) []int {

	max := -1

	for i := len(arr) - 1; i >= 0; i-- {

		arr[i], max = max, int(math.Max(float64(max), float64(arr[i])))
	}
	return arr
}
