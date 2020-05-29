package main

func countNegatives(grid [][]int) int {

	res := 0

	for _, m := range grid {
		for _, n := range m {

			if n < 0 {
				res++
			}
		}
	}

	return res
}
