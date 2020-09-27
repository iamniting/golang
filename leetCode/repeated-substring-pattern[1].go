package main

import "strings"

func repeatedSubstringPattern(s string) bool {

	for i := 0; i < len(s)/2; i++ {
		substr := s[:i+1]
		// if s can be divided in len(substr) && if s == substr * len(s) / len(substr)
		if len(s)%len(substr) == 0 && strings.Repeat(substr, len(s)/len(substr)) == s {
			return true
		}
	}

	return false
}
