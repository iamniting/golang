package main

import "unicode"

func reverseOnlyLetters(S string) string {

	res := []byte(S)
	i, j := 0, len(S)-1

	for i <= j {
		if !unicode.IsLetter(rune(S[i])) {
			i++
			continue
		}
		if !unicode.IsLetter(rune(S[j])) {
			j--
			continue
		}

		res[i], res[j] = S[j], S[i]
		i++
		j--
	}
	return string(res)
}
