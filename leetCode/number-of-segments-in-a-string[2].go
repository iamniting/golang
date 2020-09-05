package main

func countSegments(s string) int {

	var res int
	prev := ' '

	for _, c := range s {
		if prev == ' ' && c != ' ' {
			res++
		}
		prev = c
	}

	return res
}
