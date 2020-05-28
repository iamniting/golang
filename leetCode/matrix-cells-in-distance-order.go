package main

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	res := make([][]int, 0)
	queue := make([][]int, 0)
	queue = append(queue, []int{r0, c0})

	visited := make([][]bool, R)
	for i := range visited {
		visited[i] = make([]bool, C)
	}

	for len(queue) > 0 {

		r, c := queue[0][0], queue[0][1]
		queue = queue[1:]

		// check if visited continue
		if visited[r][c] {
			continue
		}

		// push to res
		res = append(res, []int{r, c})

		// push nearby points to queue
		if r-1 >= 0 {
			queue = append(queue, []int{r - 1, c})
		}
		if r+1 < R {
			queue = append(queue, []int{r + 1, c})
		}
		if c-1 >= 0 {
			queue = append(queue, []int{r, c - 1})
		}
		if c+1 < C {
			queue = append(queue, []int{r, c + 1})
		}

		// mark as visited
		visited[r][c] = true
	}

	return res
}
