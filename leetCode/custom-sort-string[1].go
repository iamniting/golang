package main

import "strings"

func customSortString(S string, T string) string {

	var output string

	for _, c := range S {

		if strings.Contains(T, string(c)) {

			count := strings.Count(T, string(c))
			output += strings.Repeat(string(c), count)
			T = strings.Replace(T, string(c), "", -1)
		}
	}

	return output + T
}
