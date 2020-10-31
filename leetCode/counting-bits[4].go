package main

func countBits(num int) []int {

	res := make([]int, num+1)
	offset := 1

	// Divide the numbers in ranges like [2-3], [4-7], [8-15] and so on. And try to generate new range from previous.
	for i := 1; i <= num; i++ {
		if offset*2 == i {
			offset *= 2
		}
		res[i] = res[i-offset] + 1
	}

	return res
}
