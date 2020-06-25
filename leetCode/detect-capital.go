package main

import "unicode"

func detectCapitalUse(word string) bool {
	var upper, lower, uIndex int

	for i, c := range word {
		if unicode.IsUpper(c) {
			upper++
			uIndex = i
		} else {
			lower++
		}
	}

	if upper != 0 && lower == 0 {
		return true
	} else if upper == 0 && lower != 0 {
		return true
	} else if upper == 1 && uIndex == 0 && lower != 0 {
		return true
	}
	return false
}
