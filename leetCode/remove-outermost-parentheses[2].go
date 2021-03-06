package main

import "bytes"

func removeOuterParentheses(S string) string {
	index := 0
	var res bytes.Buffer

	for i, c := range S {
		if c == '(' {
			index++
			if index > 1 {
				res.WriteByte(S[i])
			}
		} else {
			index--
			if index > 0 {
				res.WriteByte(S[i])
			}
		}
	}
	return res.String()
}
