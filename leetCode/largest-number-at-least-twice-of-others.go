package main

func dominantIndex(nums []int) int {

	var maxIndex int

	for i := range nums {
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
	}

	for i := range nums {
		if i != maxIndex && nums[i]*2 > nums[maxIndex] {
			return -1
		}
	}

	return maxIndex
}
