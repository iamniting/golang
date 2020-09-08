package main

func diagonalSum(mat [][]int) int {

	sum, l, j := 0, len(mat), len(mat)-1

	for i := range mat {
		// left to right diagonal
		sum += mat[i][i]
		// right to left diagonal
		sum += mat[i][j]
		j--
	}

	if len(mat)&1 == 1 {
		sum -= mat[l/2][l/2]
	}

	return sum
}
