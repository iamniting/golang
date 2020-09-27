package main

import "strings"

func reorderSpaces(text string) string {

	totalSpaces := strings.Count(text, " ")

	if totalSpaces == 0 {
		return text
	}

	words := strings.Fields(text)

	if len(words) == 1 {
		return words[0] + strings.Repeat(" ", totalSpaces)
	}

	requiredSpaces := totalSpaces / (len(words) - 1)
	res := strings.Join(words, strings.Repeat(" ", requiredSpaces))

	if len(res) != len(text) {
		return res + strings.Repeat(" ", len(text)-len(res))
	}

	return res
}
