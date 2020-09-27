package main

import "strings"

func repeatedSubstringPattern(s string) bool {

	/*
		(s + s)[1:2*len(s)-1] will have s if s itself have repeated substrings
		eg.
		abcabc
		a(bcabcabcab)c have abcabc

		eg.
		abc
		a(bcab)c does not have abc

	*/

	return strings.Contains((s + s)[1:2*len(s)-1], s)
}
