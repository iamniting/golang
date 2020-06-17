package main

import "math"

func findMaxConsecutiveOnes(nums []int) int {
	res, count := 0, 0

	for _, n := range nums {
		if n == 0 {
			count = 0
		} else {
			count++
		}
		res = int(math.Max(float64(res), float64(count)))
	}
	return res
}
