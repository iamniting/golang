package main

import "sort"

func findUnsortedSubarray(nums []int) int {

	arr := make([]int, len(nums))
	copy(arr, nums)
	sort.Ints(arr)

	min, max := -1, -1

	for i := range nums {
		if nums[i] != arr[i] {
			if min == -1 {
				min = i
			}
			max = i
		}
	}

	if min == -1 && max == -1 {
		return 0
	}

	return max - min + 1
}
