package main

import "sort"

func frequencySort(nums []int) []int {

	freq := map[int]int{}
	type valWithFreq struct{ val, freq int }
	var arr []valWithFreq
	var index int

	for _, n := range nums {
		freq[n]++
	}

	for v, f := range freq {
		arr = append(arr, valWithFreq{v, f})
	}

	sort.Slice(arr, func(i, j int) bool {
		if arr[i].freq == arr[j].freq {
			return arr[i].val > arr[j].val
		}
		return arr[i].freq < arr[j].freq
	})

	for _, obj := range arr {
		for i := 0; i < obj.freq; i++ {
			nums[index] = obj.val
			index++
		}
	}

	return nums
}
