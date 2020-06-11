package main

import "unicode"

func shortestCompletingWord(licensePlate string, words []string) string {
	m := map[rune]int{}

	for _, c := range licensePlate {
		if unicode.IsLetter(c) {
			m[unicode.ToLower(c)]++
		}
	}

	resIndx, minLen := 0, 16

	for i, str := range words {

		if len(str) >= minLen {
			continue
		}

		tmp := map[rune]int{}
		for _, c := range str {
			tmp[unicode.ToLower(c)]++
		}

		equal := 0
		flag := true
		for k := range m {
			if _, ok := tmp[k]; !ok {
				flag = false
				break
			}
			if m[k] <= tmp[k] {
				equal++
			}
		}

		if flag && equal == len(m) {
			resIndx = i
			minLen = len(str)
		}
	}

	return words[resIndx]
}
