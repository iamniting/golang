package main

import "strings"

func toGoatLatin(S string) string {
	isVowel := func(c byte) bool {
		for _, v := range []byte{
			'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'} {
			if v == c {
				return true
			}
		}
		return false
	}

	res := []string{}
	count := 1
	for _, word := range strings.Fields(S) {

		if isVowel(word[0]) {
			word = word + "ma"
		} else {
			word = word[1:] + word[0:1] + "ma"
		}

		res = append(res, word+strings.Repeat("a", count))
		count++
	}
	return strings.Join(res, " ")
}
