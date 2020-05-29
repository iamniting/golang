package main

func judgeCircle(moves string) bool {
	UD, LR := 0, 0

	for _, move := range moves {
		if move == 'U' {
			UD++
		} else if move == 'D' {
			UD--
		} else if move == 'L' {
			LR++
		} else if move == 'R' {
			LR--
		}
	}

	return UD == 0 && LR == 0
}
