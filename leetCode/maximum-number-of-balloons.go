package main

import "math"

func maxNumberOfBalloons(text string) int {

	freq := make([]int, 26)

	for _, c := range text {
		freq[c-'a']++
	}

	res := math.MaxInt32
	for _, c := range []byte{'b', 'a', 'l', 'o', 'n'} {
		if c == 'l' || c == 'o' {
			freq[c-'a'] /= 2
		}

		res = int(math.Min(float64(res), float64(freq[c-'a'])))
	}
	return res
}
