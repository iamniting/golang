package main

func rotate(nums []int, k int) {

	l := len(nums)
	k = k % l

	// get rotated array in tmp var
	tmp := append(nums[l-k:], nums[:l-k]...)
	copy(nums, tmp)
}
