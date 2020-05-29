package main

func oddCells(n int, m int, indices [][]int) int {
	row, col := make([]int, n), make([]int, m)
	oddr, oddc := 0, 0

	for _, index := range indices {
		row[index[0]]++
		col[index[1]]++
	}

	for _, v := range row {
		oddr += v & 1
	}
	for _, v := range col {
		oddc += v & 1
	}
	return oddr*(m-oddc) + oddc*(n-oddr)
}
