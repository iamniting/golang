package main

func pivotIndex(nums []int) int {
	var rsum, lsum int

	for _, n := range nums {
		rsum += n
	}

	for i, n := range nums {

		if rsum-lsum-n == lsum {
			return i
		}
		lsum += n
	}

	return -1
}
