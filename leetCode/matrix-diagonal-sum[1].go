package main

func diagonalSum(mat [][]int) int {

	sum, l := 0, len(mat)

	for i := range mat {
		// left to right diagonal
		sum += mat[i][i]
		// right to left diagonal
		sum += mat[i][l-i-1]
	}

	if l&1 == 1 {
		return sum - mat[l/2][l/2]
	}

	return sum
}
