package main

func tictactoe(moves [][]int) string {
	row, col, dig := make([]int, 3), make([]int, 3), make([]int, 2)

	for i, mv := range moves {
		r, c := mv[0], mv[1]
		if i&1 == 0 {
			row[r]++
			col[c]++
			if r == c {
				dig[0]++
			}
			if r+c == 2 {
				dig[1]++
			}
			if row[r] == 3 || col[c] == 3 || dig[0] == 3 || dig[1] == 3 {
				return "A"
			}
		} else {
			row[r]--
			col[c]--
			if r == c {
				dig[0]--
			}
			if r+c == 2 {
				dig[1]--
			}
			if row[r] == -3 || col[c] == -3 || dig[0] == -3 || dig[1] == -3 {
				return "B"
			}
		}
	}

	if len(moves) == 9 {
		return "Draw"
	}
	return "Pending"
}
