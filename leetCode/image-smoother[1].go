package main

func imageSmoother(M [][]int) [][]int {
	m, n := len(M), len(M[0])
	res := make([][]int, m)

	for i, row := range M {
		res[i] = make([]int, n)

		for j := range row {
			var count int

			// iterate through neighbours
			for r := -1; r < 2; r++ {
				for c := -1; c < 2; c++ {
					x, y := i+r, j+c
					if x >= 0 && x < m && y >= 0 && y < n {
						res[i][j] += M[x][y]
						count++
					}
				}
			}
			res[i][j] /= count
		}
	}
	return res
}
