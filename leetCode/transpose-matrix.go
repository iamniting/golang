package main

func transpose(A [][]int) [][]int {

	res := make([][]int, len(A[0]))

	for i := range A[0] {
		res[i] = make([]int, len(A))
	}

	for i, row := range A {

		for j := range row {

			res[j][i] = A[i][j]
		}
	}

	return res
}
