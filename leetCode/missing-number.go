package main

func missingNumber(nums []int) int {

	missing := len(nums)

	for i, n := range nums {
		missing ^= i ^ n
	}

	return missing
}
