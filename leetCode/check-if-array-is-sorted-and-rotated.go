package main

func check(nums []int) bool {

	var isSorted int

	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			isSorted++
		}
	}

	return isSorted == 0 || isSorted == 1 && nums[0] >= nums[len(nums)-1]
}
