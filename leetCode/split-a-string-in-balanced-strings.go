package main

func balancedStringSplit(s string) int {
	res, bal := 0, 0

	for _, c := range s {
		if c == 'L' {
			bal++
		} else {
			bal--
		}

		if bal == 0 {
			res++
		}
	}
	return res
}
