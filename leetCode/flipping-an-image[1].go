package main

func flipAndInvertImage(A [][]int) [][]int {
	for _, row := range A {
		i, j := 0, len(row)-1

		for i <= j {
			row[i], row[j] = row[j]^1, row[i]^1
			i++
			j--
		}
	}
	return A
}
