package main

func shiftGrid(grid [][]int, k int) [][]int {
	k = k % (len(grid) * len(grid[0]))

	// create a tmp arr
	tmp := []int{}
	for _, row := range grid {
		tmp = append(tmp, row...)
	}
	tmp = append(tmp[len(tmp)-k:], tmp[:len(tmp)-k]...)

	// generate grid from tmp arr
	index := 0
	for i, row := range grid {
		for j := range row {
			grid[i][j] = tmp[index]
			index++
		}
	}

	return grid
}
