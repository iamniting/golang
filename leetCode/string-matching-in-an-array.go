package main

import "strings"

func stringMatching(words []string) []string {

	resMap := make(map[string]bool)
	res := []string{}

	for i, word := range words {
		for j, wrd := range words {
			if i != j && strings.Contains(wrd, word) {
				resMap[word] = true
			}
		}
	}

	for key := range resMap {
		res = append(res, key)
	}

	return res
}
