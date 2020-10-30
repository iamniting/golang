package main

func countBits(num int) []int {

	res := make([]int, num+1)

	for i := 0; i <= num; i++ {
		if i&1 == 1 {
			res[i] = res[i-1] + 1
			continue
		}

		for j := i; j > 0; j >>= 1 {
			res[i] += j & 1
		}
	}

	return res
}
