package main

func islandPerimeter(grid [][]int) int {
	res, m, n := 0, len(grid), len(grid[0])

	for i, row := range grid {
		for j, col := range row {
			if col == 0 {
				continue
			}

			res += 4
			// check up down left right
			if i > 0 && grid[i-1][j] == 1 {
				res--
			}
			if i < m-1 && grid[i+1][j] == 1 {
				res--
			}
			if j > 0 && grid[i][j-1] == 1 {
				res--
			}
			if j < n-1 && grid[i][j+1] == 1 {
				res--
			}
		}
	}
	return res
}
