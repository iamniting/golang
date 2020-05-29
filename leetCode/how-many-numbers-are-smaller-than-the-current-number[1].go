package main

func smallerNumbersThanCurrent(nums []int) []int {

	occurance := make([]int, 101)
	res := make([]int, len(nums))

	// get occurance of each num
	for _, n := range nums {
		occurance[n]++
	}

	// store count of lesser nums
	for i := 1; i < 101; i++ {
		occurance[i] = occurance[i] + occurance[i-1]
	}

	// get result
	for i, n := range nums {

		if n == 0 {
			res[i] = 0
		} else {
			res[i] = occurance[n-1]
		}
	}

	return res
}
