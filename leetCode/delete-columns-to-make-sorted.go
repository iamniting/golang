package main

func minDeletionSize(A []string) int {
	if A == nil {
		return 0
	}

	m, n, res := len(A), len(A[0]), 0

	for c := 0; c < n; c++ {
		for r := 0; r < m-1; r++ {
			if A[r][c] > A[r+1][c] {
				res++
				break
			}
		}
	}
	return res
}
