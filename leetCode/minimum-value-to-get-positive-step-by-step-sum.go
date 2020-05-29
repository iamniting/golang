package main

import "math"

func minStartValue(nums []int) int {
	res := 0

	// run utill does not get ans
	for res = 1; res <= math.MaxInt32; res++ {
		sum, flag := res, true

		for _, n := range nums {
			sum += n
			if sum <= 0 {
				flag = false
				break
			}
		}
		// if does not went into sum <= 0 state then return
		if flag {
			return res
		}
	}
	return res
}
