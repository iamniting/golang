package main

import "math"

func findMaxAverage(nums []int, k int) float64 {

	getSum := func(arr []int) (sum float64) {
		for _, n := range arr {
			sum += float64(n)
		}
		return sum
	}

	res := float64(math.MinInt32)

	for i := 0; i <= len(nums)-k; i++ {
		avg := getSum(nums[i:i+k]) / float64(k)
		res = math.Max(res, avg)
	}

	return res
}
