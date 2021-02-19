package main

import "math"

func mergeAlternately(word1 string, word2 string) string {

	var res string
	l := int(math.Min(float64(len(word1)), float64(len(word2))))

	for i := 0; i < l+1; i++ {
		if len(word1) == i {
			res += word2[i:]
		} else if len(word2) == i {
			res += word1[i:]
		} else {
			res += string(word1[i]) + string(word2[i])
		}
	}

	return res
}
