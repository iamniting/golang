package main

import "sort"

func numSmallerByFrequency(queries []string, words []string) []int {
	getFreq := func(s string) int {
		freq := make([]int, 26)
		for i := range s {
			freq[s[i]-'a']++
		}

		for _, n := range freq {
			if n != 0 {
				return n
			}
		}
		return 0
	}

	wordsFreq := make([]int, len(words))
	res := make([]int, len(queries))

	for i, s := range words {
		wordsFreq[i] = getFreq(s)
	}

	sort.Ints(wordsFreq)

	for i, q := range queries {
		freq := getFreq(q)
		res[i] = len(wordsFreq) - sort.SearchInts(wordsFreq, freq+1)
	}

	return res
}
