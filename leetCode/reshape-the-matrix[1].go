package main

func matrixReshape(nums [][]int, r int, c int) [][]int {
	m, n := len(nums), len(nums[0])
	if r*c != m*n {
		return nums
	}

	res := make([][]int, r)
	for i := range res {
		res[i] = make([]int, c)
	}

	for i := 0; i < r*c; i++ {
		res[i/c][i%c] = nums[i/n][i%n]
	}
	return res
}
