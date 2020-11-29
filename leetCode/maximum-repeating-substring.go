package main

import "math"

func maxRepeating(sequence string, word string) int {

	var res, k int

	for i := 0; i < len(sequence); i++ {
		if i+len(word) <= len(sequence) && sequence[i:i+len(word)] == word {
			i += len(word) - 1
			k++
			res = int(math.Max(float64(res), float64(k)))
		} else {
			k = 0
		}
	}

	return res
}
