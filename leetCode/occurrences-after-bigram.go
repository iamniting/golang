package main

import "strings"

func findOcurrences(text string, first string, second string) []string {

	tmp := strings.Fields(text)
	ans := []string{}

	for i := 0; i < len(tmp)-2; i++ {
		if tmp[i] == first && tmp[i+1] == second {
			ans = append(ans, tmp[i+2])
		}
	}

	return ans
}
