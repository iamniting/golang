package main

import "math"

func hasGroupsSizeX(deck []int) bool {

	m := map[int]int{}

	// count freq
	for _, n := range deck {
		m[n]++
	}

	min := math.MaxInt32
	odd := false

	// get minimum
	// check if odd presents
	// retun false if v < 2
	for _, v := range m {
		min = int(math.Min(float64(min), float64(v)))
		if v&1 == 1 {
			odd = true
		}
		if v < 2 {
			return false
		}
	}

	// if no odd num presents we can split cards in group of 2
	if !odd {
		return true
	}

	// find common divisor if presents return true
	for i := 3; i <= min; i++ {
		var flag bool
		for _, v := range m {
			if v%i != 0 {
				flag = true
				break
			}
		}

		if !flag {
			return true
		}
	}

	return false
}
