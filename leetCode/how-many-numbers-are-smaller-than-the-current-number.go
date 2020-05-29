package main

func smallerNumbersThanCurrent(nums []int) []int {

	res := make([]int, len(nums))

	for i := range nums {

		for j := range nums {

			if nums[i] > nums[j] {
				res[i]++
			}
		}
	}

	return res
}
