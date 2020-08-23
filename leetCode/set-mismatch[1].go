package main

import "math"

func findErrorNums(nums []int) []int {

	var duplicate, missing int

	// get duplicate
	for _, n := range nums {
		if nums[int(math.Abs(float64(n)))-1] < 0 {
			// get the duplicate if num is already negative
			duplicate = int(math.Abs(float64(n)))
		} else {
			// mark the n - 1 index as negative
			nums[int(math.Abs(float64(n)))-1] *= -1
		}
	}

	// get missing
	for i, n := range nums {
		// if n is not -ve it means that num is not marked and found
		if n > 0 {
			missing = i + 1
		}
	}

	return []int{duplicate, missing}
}
