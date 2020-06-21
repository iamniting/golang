package main

import "math"

func isAlienSorted(words []string, order string) bool {
	slice := make([]int, 26)

	for i, c := range order {
		slice[c-'a'] = i
	}

	for i := 1; i < len(words); i++ {

		w1, w2 := words[i-1], words[i]
		minLen := int(math.Min(float64(len(w1)), float64(len(w2))))

		for j := 0; j < minLen; j++ {
			if slice[w1[j]-'a'] > slice[w2[j]-'a'] {
				return false
			}
			if slice[w1[j]-'a'] < slice[w2[j]-'a'] {
				break
			}
		}

		if w1[minLen-1] == w2[minLen-1] && len(w1) > len(w2) {
			return false
		}
	}

	return true
}
