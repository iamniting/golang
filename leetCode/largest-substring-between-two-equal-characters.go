package main

import "math"

func maxLengthBetweenEqualCharacters(s string) int {

	res := float64(-1)
	var index [26]int

	for i := range index {
		index[i] = -1
	}

	for i, c := range s {
		if index[c-'a'] == -1 {
			index[c-'a'] = i
			continue
		}

		res = math.Max(res, float64(i-index[c-'a']-1))
	}

	return int(res)
}
