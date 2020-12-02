package main

import "math"

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {

	res := make([]bool, len(l))
	for i := range l {
		res[i] = isArithmetic(nums[l[i] : r[i]+1])
	}
	return res
}

func isArithmetic(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}

	set := make(map[int]bool)
	min, max := math.MaxInt32, math.MinInt32

	for _, n := range nums {
		min = int(math.Min(float64(min), float64(n)))
		max = int(math.Max(float64(max), float64(n)))
		set[n] = true
	}

	// if can not divide return false
	// eg. 3,4,5,6,8 will be 8-3 = 5; 5 / 4 = 1.25
	if (max-min)%(len(nums)-1) != 0 {
		return false
	}

	interval := (max - min) / (len(nums) - 1)

	// if all required nums are not present in set return false
	for i := 1; i < len(nums); i++ {
		if !set[min+i*interval] {
			return false
		}
	}
	return true
}
