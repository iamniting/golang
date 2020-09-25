package main

func gardenNoAdj(N int, paths [][]int) []int {

	graph := make([][]int, N)
	res := make([]int, N)

	// create bidirectional graph
	for _, p := range paths {
		x, y := p[0]-1, p[1]-1
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}

	for i, row := range graph {

		color := [5]bool{}

		// check what all colors present in adjacent gardens
		for _, col := range row {
			color[res[col]] = true
		}

		// set color which is not present in adjacent gardens
		for j := 1; j < 5; j++ {
			if !color[j] {
				res[i] = j
				break
			}
		}
	}

	return res
}
