package main

import "sort"

func minSubsequence(nums []int) []int {
	sort.Slice(nums, func(a, b int) bool { return nums[a] > nums[b] })

	sum := 0
	for _, n := range nums {
		sum += n
	}

	resSum := 0
	res := []int{}

	for _, n := range nums {
		resSum += n
		res = append(res, n)
		if resSum > sum-resSum {
			break
		}
	}
	return res
}
