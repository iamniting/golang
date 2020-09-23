package main

import "sort"

func longestWord(words []string) string {

	var res string
	set := make(map[string]bool, len(words))

	// sort the strings
	sort.Strings(words)

	for _, word := range words {
		// if word except last char does not exist then there is no need
		// to add that in the set as we will be not able to make words with it
		// e.g. a, ab, abc if ab does not exist it does not make any sense to
		// add abc as we will be not able to make abcd out of that
		if len(word) > 1 && !set[word[:len(word)-1]] {
			continue
		}

		set[word] = true
		if len(word) > len(res) {
			res = word
		}
	}
	return res
}
