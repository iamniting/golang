package main

import "strings"

func customSortString(S string, T string) string {

	var res string
	var freqT [26]int

	for _, c := range T {
		freqT[c-'a']++
	}

	for _, c := range S {
		res += strings.Repeat(string(c), freqT[c-'a'])
		freqT[c-'a'] = 0
	}

	for i, n := range freqT {
		if n != 0 {
			res += strings.Repeat(string(i+'a'), n)
		}
	}

	return res
}
