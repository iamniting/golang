package main

import "math"

func findDisappearedNumbers(nums []int) []int {

	for _, n := range nums {
		n = int(math.Abs(float64(n)))
		nums[n-1] = -int(math.Abs(float64(nums[n-1])))
	}

	res := []int{}
	for i, n := range nums {
		if n > 0 {
			res = append(res, i+1)
		}
	}

	return res
}
