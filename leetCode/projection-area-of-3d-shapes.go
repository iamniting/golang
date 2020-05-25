package main

import "math"

func projectionArea(grid [][]int) int {

	res := 0

	for i, row := range grid {

		rmax := 0.0
		cmax := 0.0

		for j, col := range row {

			if col != 0 {
				res++
			}
			rmax = math.Max(rmax, float64(grid[i][j]))
			cmax = math.Max(cmax, float64(grid[j][i]))

		}
		res += int(rmax) + int(cmax)

	}
	return res
}

