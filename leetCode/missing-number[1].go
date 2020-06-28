package main

func missingNumber(nums []int) int {

	sum := (len(nums) * (len(nums) + 1)) / 2

	var actualsum int
	for _, n := range nums {
		actualsum += n
	}

	return sum - actualsum
}
