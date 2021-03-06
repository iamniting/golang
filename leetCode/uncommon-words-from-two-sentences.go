package main

import "strings"

func uncommonFromSentences(A string, B string) []string {

	res := []string{}
	m := make(map[string]int)

	// store words and occurrence
	for _, word := range strings.Fields(A + " " + B) {
		m[word]++
	}

	// if occurrence is 1 append it into res
	for k, v := range m {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}
