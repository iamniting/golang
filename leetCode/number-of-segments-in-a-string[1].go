package main

func countSegments(s string) int {

	var res int
	var isFirst bool

	for _, c := range s {
		if !isFirst && c != ' ' {
			res++
			isFirst = true
		} else if c == ' ' {
			isFirst = false
		}
	}

	return res
}
