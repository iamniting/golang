package main

import "strconv"

func freqAlphabets(s string) string {
	res := ""

	for i := len(s) - 1; i >= 0; i-- {
		num := 0
		if s[i] == '#' {
			num, _ = strconv.Atoi(s[i-2 : i])
			i = i - 2
		} else {
			num, _ = strconv.Atoi(s[i : i+1])
		}
		res = string(byte(num+96)) + res
	}
	return res
}
