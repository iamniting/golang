package main

import "unicode"

func shortestCompletingWord(licensePlate string, words []string) string {

	count := func(word string) []int {
		res := make([]int, 26)

		for _, c := range word {
			if unicode.IsLetter(c) {
				res[unicode.ToLower(c)-'a']++
			}
		}
		return res
	}

	isValid := func(freq []int, word string) bool {
		tmp := count(word)
		for i := range freq {
			if freq[i] > tmp[i] {
				return false
			}
		}
		return true
	}

	res, minlen := "", 16
	freq := count(licensePlate)
	for _, word := range words {
		if minlen > len(word) && isValid(freq, word) {
			res = word
			minlen = len(word)
		}
	}

	return res
}
