package main

func checkRecord(s string) bool {

	A, L1, L2 := 0, -2, -2

	for i, c := range s {
		if c == 'A' {
			A++
			if A > 1 {
				return false
			}

		} else if c == 'L' {
			if L2-L1 == 1 && i-L2 == 1 {
				return false
			}
			L1, L2 = L2, i
		}
	}

	return true
}
