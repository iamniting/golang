package main

func countConsistentStrings(allowed string, words []string) int {

	var res int
	var slice [26]bool

	for _, c := range allowed {
		slice[c-'a'] = true
	}

	for _, word := range words {
		res++
		for _, c := range word {
			if !slice[c-'a'] {
				res--
				break
			}
		}
	}
	return res
}
