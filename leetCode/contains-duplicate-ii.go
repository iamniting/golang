package main

import "math"

func containsNearbyDuplicate(nums []int, k int) bool {

	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] && int(math.Abs(float64(j-i))) <= k {
				return true
			}
		}
	}

	return false
}
