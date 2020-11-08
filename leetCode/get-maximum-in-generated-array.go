package main

import "math"

func getMaximumGenerated(n int) int {

	if n == 0 || n == 1 {
		return n
	}

	var res int
	slice := make([]int, n+1)
	slice[1] = 1

	for i := 2; i <= n; i++ {
		if i&1 == 0 {
			slice[i] = slice[i/2]
		} else {
			slice[i] = slice[i/2] + slice[i/2+1]
		}

		res = int(math.Max(float64(res), float64(slice[i])))
	}

	return res
}
