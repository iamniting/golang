package main

import "unicode"

func numRookCaptures(board [][]byte) int {

	res := 0

	for i, row := range board {
		for j, col := range row {

			if col == 'R' {
				for k := j - 1; k >= 0; k-- {
					if unicode.IsUpper(rune(row[k])) {
						break
					}
					if row[k] == 'p' {
						res++
						break
					}
				}
				for k := j + 1; k < 8; k++ {
					if unicode.IsUpper(rune(row[k])) {
						break
					}
					if row[k] == 'p' {
						res++
						break
					}
				}
				for k := i - 1; k >= 0; k-- {
					if unicode.IsUpper(rune(board[k][j])) {
						break
					}
					if board[k][j] == 'p' {
						res++
						break
					}
				}
				for k := i + 1; k < 8; k++ {
					if unicode.IsUpper(rune(board[k][j])) {
						break
					}
					if board[k][j] == 'p' {
						res++
						break
					}
				}
				return res
			}
		}
	}
	return 0
}
