package main

func numberOfMatches(n int) int {

	var res int

	for n > 1 {
		res += n / 2
		n = (n / 2) + (n & 1)
	}
	return res
}
