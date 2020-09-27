package main

import (
	"strings"
	"unicode"
)

func mostCommonWord(paragraph string, banned []string) string {

	ban := map[string]bool{}
	m := map[string]int{}

	for _, s := range banned {
		ban[s] = true
	}

	res := ""

	// remove all punctuation symbols !?',;. in FielsFunc
	for _, s := range strings.FieldsFunc(paragraph, func(c rune) bool {
		return !unicode.IsLetter(c)
	}) {

		s = strings.ToLower(s)
		m[s]++

		if _, ok := ban[s]; !ok && m[s] > m[res] {
			res = s
		}
	}

	return res
}
