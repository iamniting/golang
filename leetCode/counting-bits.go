package main

func countBits(num int) []int {

	res := make([]int, num+1)

	for i := 0; i <= num; i++ {

		for j := i; j > 0; j >>= 1 {
			res[i] += j & 1
		}
	}

	return res
}
