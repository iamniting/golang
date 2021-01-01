package main

func minFlips(target string) int {

	var res int
	next := byte('0')

	for i := range target {
		if target[i] != next {
			res++
			// flip the next 0 to 1 and 1 to 0
			next = 97 - next
		}
	}

	return res
}
