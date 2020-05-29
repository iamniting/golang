package main

import "math"

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {

	res := 0

	for _, i := range arr1 {

		flag := true
		for _, j := range arr2 {
			if math.Abs(float64(i-j)) <= float64(d) {
				flag = false
			}
		}
		if flag {
			res++
		}
	}

	return res
}
