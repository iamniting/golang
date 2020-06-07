package main

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

	for i, q := range queries {
		freq := getFreq(q)

		for _, n := range wordsFreq {
			if n > freq {
				res[i]++
			}
		}
	}

	return res
}
