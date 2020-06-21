package main

import "unicode"

func reformat(s string) string {

	// build a string
	build := func(s1, s2 []rune) string {
		res := []rune{}
		for i := range s2 {
			res = append(res, s1[i])
			res = append(res, s2[i])
		}
		if len(s1) != len(s2) {
			res = append(res, s1[len(s1)-1])
		}
		return string(res)
	}

	// get digits and letters in a slice
	var digits, letters []rune
	for _, c := range s {
		if unicode.IsDigit(c) {
			digits = append(digits, c)
		} else {
			letters = append(letters, c)
		}
	}

	// if combination will not be possible return empty string
	if len(digits)-len(letters) > 1 || len(letters)-len(digits) > 1 {
		return ""
	}

	if len(digits) > len(letters) {
		return build(digits, letters)
	}
	return build(letters, digits)
}
