package main

func checkPossibility(nums []int) bool {

	var flag bool

	for i := 0; i < len(nums)-1; i++ {

		if nums[i] > nums[i+1] {

			if flag {
				return false
			}

			if i == 0 || nums[i-1] <= nums[i+1] {
				nums[i] = nums[i+1]
			} else {
				nums[i+1] = nums[i]
			}
			flag = true
		}
	}

	return true
}
