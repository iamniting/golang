package main

import "strings"

func wordPattern(pattern string, str string) bool {

	words := strings.Fields(str)
	m1 := map[byte]string{}
	m2 := map[string]byte{}

	if len(pattern) != len(words) {
		return false
	}

	for i := range pattern {
		p, w := pattern[i], words[i]

		if val, ok := m1[p]; ok && val != w {
			return false
		}

		if val, ok := m2[w]; ok && val != p {
			return false
		}

		m1[p] = w
		m2[w] = p
	}

	return true
}
