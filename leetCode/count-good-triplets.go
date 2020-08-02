package main

import "math"

func countGoodTriplets(arr []int, a int, b int, c int) int {

	var count int

	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			for k := j + 1; k < len(arr); k++ {
				if int(math.Abs(float64(arr[i]-arr[j]))) <= a &&
					int(math.Abs(float64(arr[j]-arr[k]))) <= b &&
					int(math.Abs(float64(arr[k]-arr[i]))) <= c {
					count++
				}
			}
		}
	}

	return count
}
