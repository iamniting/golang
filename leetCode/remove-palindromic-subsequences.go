package main

func removePalindromeSub(s string) int {
	isPalindrome := func(str string) bool {
		for i := 0; i < len(s)/2; i++ {
			if s[i] != s[len(s)-1-i] {
				return false
			}
		}
		return true
	}

	if len(s) == 0 {
		return 0
	}
	if isPalindrome(s) {
		return 1
	}

	/*
	   You need to know the difference between subarray and subsequence.
	   Subarray need to be consecutive。
	   Subsequence don't have to be consecutive.

	   Explanation
	   We can delete all characters 'a' in the 1st operation,
	   and then all characters 'b' in the 2nd operation.
	   So return 2 in this case
	*/
	return 2
}
