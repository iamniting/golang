package main

import "math"

func findUnsortedSubarray(nums []int) int {

	/*
		The idea behind this method is that the correct position of the
		minimum element in the unsorted subarray helps to determine the
		required left boundary. Similarly, the correct position of the
		maximum element in the unsorted subarray helps to determine the
		required right boundary.

		Thus, firstly we need to determine when the correctly sorted
		array goes wrong. We keep a track of this by observing rising
		slope starting from the beginning of the array. Whenever the
		slope falls, we know that the unsorted array has surely started.
		Thus, now we determine the minimum element found till the end
		of the array nums, given by min.

		Similarly, we scan the array nums in the reverse order and when
		the slope becomes rising instead of falling, we start looking
		for the maximum element till we reach the beginning of the
		array, given by max.

		Then, we traverse over nums and determine the correct position
		of min and max by comparing these elements with the other array
		elements. e.g. To determine the correct position of min, we
		know the initial portion of nums is already sorted. Thus, we
		need to find the first element which is just larger than min.
		Similarly, for max's position, we need to find the first element
		which is just smaller than maxmax searching in numsnums backwards.
	*/

	min, max := math.MaxInt32, math.MinInt32

	flag := false
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			flag = true
		}
		if flag {
			min = int(math.Min(float64(min), float64(nums[i])))
		}
	}

	flag = false
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] > nums[i+1] {
			flag = true
		}
		if flag {
			max = int(math.Max(float64(max), float64(nums[i])))
		}
	}

	var l, r int

	for l = 0; l < len(nums); l++ {
		if min < nums[l] {
			break
		}
	}

	for r = len(nums) - 1; r >= 0; r-- {
		if max > nums[r] {
			break
		}
	}

	if r-l < 0 {
		return 0
	}

	return r - l + 1
}
