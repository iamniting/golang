package main

import (
	"sort"
	"strings"
	"unicode"
)

func reorderLogFiles(logs []string) []string {
	let, dig := []string{}, []string{}

	for _, log := range logs {
		str := strings.Fields(log)
		if unicode.IsLetter(rune(str[1][0])) {
			let = append(let, log)
		} else {
			dig = append(dig, log)
		}
	}

	sort.SliceStable(let, func(i, j int) bool {

		iSpace := strings.Index(let[i], " ")
		jSpace := strings.Index(let[j], " ")

		iLog := let[i][iSpace:]
		jLog := let[j][jSpace:]

		if iLog == jLog {
			return let[i] < let[j]
		}
		return iLog < jLog
	})

	return append(let, dig...)
}
