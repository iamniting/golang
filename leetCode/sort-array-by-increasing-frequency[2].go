package main

import "sort"

func frequencySort(nums []int) []int {

	count := map[int]int{}
	var freq [101][]int
	var index int

	for _, n := range nums {
		count[n]++
	}

	for v, f := range count {
		freq[f] = append(freq[f], v)
	}

	for f, values := range freq {
		sort.Ints(values)
		for v := len(values) - 1; v >= 0; v-- {
			for i := 0; i < f; i++ {
				nums[index] = values[v]
				index++
			}
		}
	}

	return nums
}
