package main

func gardenNoAdj(N int, paths [][]int) []int {

	graph := map[int][]int{}
	res := make([]int, N)

	// create bidirectional graph
	for _, p := range paths {
		x, y := p[0], p[1]
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}

	for i := 1; i <= N; i++ {

		color := [5]bool{}

		// check what all colors present in adjacent gardens
		for _, g := range graph[i] {
			color[res[g-1]] = true
		}

		// set color which is not present in adjacent gardens
		for j := 1; j < 5; j++ {
			if !color[j] {
				res[i-1] = j
				break
			}
		}
	}

	return res
}
