package main

func matrixReshape(nums [][]int, r int, c int) [][]int {

	if r*c != len(nums)*len(nums[0]) {
		return nums
	}

	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}

	i, j := 0, 0
	for _, row := range nums {
		for _, col := range row {
			res[i][j] = col
			j++
			if j == c {
				j = 0
				i++
			}
		}
	}

	return res
}
