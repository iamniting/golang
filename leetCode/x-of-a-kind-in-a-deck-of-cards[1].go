package main

func hasGroupsSizeX(deck []int) bool {

	var res int
	count := map[int]int{}

	// count freq
	for _, n := range deck {
		count[n]++
	}

	// get common divisor
	for _, v := range count {
		if v < 2 {
			return false
		}
		res = gcd(res, v)
	}

	return res > 1
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
