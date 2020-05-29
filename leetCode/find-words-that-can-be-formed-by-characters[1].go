package main

func countCharacters(words []string, chars string) int {

	frequency := [26]int{}
	for _, c := range chars {
		frequency[c-'a']++
	}

	good := func(word string, frequency [26]int) bool {
		for _, c := range word {
			frequency[c-'a']--
			if frequency[c-'a'] < 0 {
				return false
			}
		}
		return true
	}

	goodWord := 0
	for _, word := range words {
		if good(word, frequency) {
			goodWord += len(word)
		}
	}
	return goodWord
}
