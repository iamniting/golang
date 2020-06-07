package main

func shuffle(nums []int, n int) []int {

	res := []int{}
	x, y := 0, n

	for ; x < n; x++ {
		res = append(res, nums[x])
		res = append(res, nums[y])
		y++
	}

	return res
}
