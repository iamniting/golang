package main

func imageSmoother(M [][]int) [][]int {
	m, n := len(M), len(M[0])
	res := make([][]int, m)

	points := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 0}, {0, 1},
		{1, -1}, {1, 0}, {1, 1}}

	for i, row := range M {
		res[i] = make([]int, n)

		for j := range row {
			var count int

			// iterate through neighbours
			for _, p := range points {
				x, y := i+p[0], j+p[1]
				if x >= 0 && x < m && y >= 0 && y < n {
					res[i][j] += M[x][y]
					count++
				}
			}
			res[i][j] /= count
		}
	}
	return res
}
