package main

func numMovesStones(a int, b int, c int) []int {

	res := make([]int, 2)

	// sort
	if a > b && b > c {
		a, c = c, a
	} else if a > c && c > b {
		a, b, c = b, c, a
	} else if b > c && c > a {
		b, c = c, b
	} else if b > a && a > c {
		a, b, c = c, a, b
	} else if c > a && a > b {
		a, b = b, a
	}

	// minimum moves
	if c-b == 1 && b-a == 1 {
		res[0] = 0
	} else if c-b > 2 && b-a > 2 {
		res[0] = 2
	} else {
		res[0] = 1
	}

	// maximum moves
	res[1] = (c - b - 1) + (b - a - 1)

	return res
}
