package main

import "math"

func surfaceArea(grid [][]int) int {

	res, m, n := 0, len(grid), len(grid[0])

	for i, row := range grid {
		for j, col := range row {
			if col == 0 {
				continue
			}

			res += (col * 4) + 2

			// check up down left right
			if i > 0 {
				res -= int(math.Min(float64(col), float64(grid[i-1][j])))
			}
			if i < m-1 {
				res -= int(math.Min(float64(col), float64(grid[i+1][j])))
			}
			if j > 0 {
				res -= int(math.Min(float64(col), float64(grid[i][j-1])))
			}
			if j < n-1 {
				res -= int(math.Min(float64(col), float64(grid[i][j+1])))
			}
		}
	}
	return res
}
