package main

func smallerNumbersThanCurrent(nums []int) []int {

	occurrence := make([]int, 101)
	res := make([]int, len(nums))

	// get occurrence of each num
	for _, n := range nums {
		occurrence[n]++
	}

	// store count of lesser nums
	for i := 1; i < 101; i++ {
		occurrence[i] = occurrence[i] + occurrence[i-1]
	}

	// get result
	for i, n := range nums {

		if n == 0 {
			res[i] = 0
		} else {
			res[i] = occurrence[n-1]
		}
	}

	return res
}
