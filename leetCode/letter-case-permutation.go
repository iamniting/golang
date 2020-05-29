package main

import "unicode"

func letterCasePermutation(S string) []string {

	res := []string{""}

	for _, c := range S {
		// if c is numeric just add c char into all strings in res
		if unicode.IsNumber(c) {
			for i := range res {
				res[i] += string(c)
			}
			continue
		}
		// if c is alpabet add lower and upper of the char
		for i := range res {
			res = append(res, res[i]+string(unicode.ToUpper(c)))
			res[i] += string(unicode.ToLower(c))
		}
	}
	return res
}
