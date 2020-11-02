package main

import "sort"

func frequencySort(nums []int) []int {

	freq := map[int]int{}

	for _, n := range nums {
		freq[n]++
	}

	sort.Slice(nums, func(i, j int) bool {
		if freq[nums[i]] == freq[nums[j]] {
			return nums[i] > nums[j]
		}
		return freq[nums[i]] < freq[nums[j]]
	})

	return nums
}
