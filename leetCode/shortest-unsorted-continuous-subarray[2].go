package main

import "math"

func findUnsortedSubarray(nums []int) int {

	var stack []int
	minInd, maxInd := math.MaxInt32, math.MinInt32

	for i := 0; i < len(nums); i++ {
		l := len(stack)
		for l > 0 && nums[stack[l-1]] > nums[i] {
			minInd = int(math.Min(float64(minInd), float64(stack[l-1])))
			stack = stack[:l-1]
			l = len(stack)
		}
		stack = append(stack, i)
	}

	stack = []int{}

	for i := len(nums) - 1; i >= 0; i-- {
		l := len(stack)
		for l > 0 && nums[stack[l-1]] < nums[i] {
			maxInd = int(math.Max(float64(maxInd), float64(stack[l-1])))
			stack = stack[:l-1]
			l = len(stack)
		}
		stack = append(stack, i)
	}

	if minInd == math.MaxInt32 && maxInd == math.MinInt32 {
		return 0
	}

	return maxInd - minInd + 1
}
