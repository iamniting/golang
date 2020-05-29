package main

import "strings"

func generateTheString(n int) string {

	// odd num
	if n&1 == 1 {
		return strings.Repeat("a", n)
	}

	// even num
	return strings.Repeat("a", n-1) + "b"
}
