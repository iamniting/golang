package main

import "sort"

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {

	res := make([]bool, len(l))

	for i := range l {

		res[i] = true
		slice := make([]int, r[i]-l[i]+1)
		copy(slice, nums[l[i]:r[i]+1])
		sort.Ints(slice)
		diff := slice[1] - slice[0]

		for j := 2; j < len(slice); j++ {
			if slice[j]-slice[j-1] != diff {
				res[i] = false
				break
			}
		}
	}

	return res
}
