package main

func maxSubArray(nums []int) int {

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	res := ^(1 << 32) + 1

	for i := range nums {
		var sum int
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			res = max(res, sum)
		}
	}

	return res
}
