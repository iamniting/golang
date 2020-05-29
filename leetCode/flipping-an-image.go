package main

func flipAndInvertImage(A [][]int) [][]int {
	res := make([][]int, len(A))

	for i, row := range A {
		res[i] = make([]int, len(row))

		// reverse the row and flip the image
		index := len(row) - 1
		for _, num := range row {
			res[i][index] = num ^ 1
			index--
		}
	}
	return res
}
