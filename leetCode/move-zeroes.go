package main

func moveZeroes(nums []int) {
	n := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[n], nums[i] = nums[i], nums[n]
			n++
		}
	}
}
