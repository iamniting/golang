package main

func numSpecial(mat [][]int) int {

	var res int
	rows, cols := make([]int, len(mat)), make([]int, len(mat[0]))

	// count 1 in rows and cols
	for i, row := range mat {
		for j, col := range row {
			rows[i] += col
			cols[j] += col
		}
	}

	for i, row := range mat {
		for j, col := range row {
			// if exactly 1 then res++
			if col == 1 && rows[i] == 1 && cols[j] == 1 {
				res++
			}
		}
	}

	return res
}
