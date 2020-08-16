package main

func validPalindrome(s string) bool {

	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return isPal(s[i:j]) || isPal(s[i+1:j+1])
		}
	}

	return true
}

func isPal(s string) bool {

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}
