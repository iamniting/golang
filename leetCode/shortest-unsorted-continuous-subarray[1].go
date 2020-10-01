package main

import "math"

func findUnsortedSubarray(nums []int) int {

	minInd, maxInd := math.MaxInt32, math.MinInt32

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				minInd = int(math.Min(float64(minInd), float64(i)))
				maxInd = int(math.Max(float64(maxInd), float64(j)))
			}
		}
	}

	if minInd == math.MaxInt32 && maxInd == math.MinInt32 {
		return 0
	}

	return maxInd - minInd + 1
}
