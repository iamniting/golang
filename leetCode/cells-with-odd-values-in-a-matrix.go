package main

func oddCells(n int, m int, indices [][]int) int {
	// initialize 2d matrix with zero val
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
	}

	// increment row and col as per given indices
	for _, ind := range indices {
		row, col := ind[0], ind[1]
		// increment row
		for i := 0; i < m; i++ {
			matrix[row][i]++
		}
		// increment col
		for i := 0; i < n; i++ {
			matrix[i][col]++
		}
	}

	// count odds
	res := 0
	for _, row := range matrix {
		for _, col := range row {
			res += col & 1
		}
	}
	return res
}
