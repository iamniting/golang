package main

func repeatedSubstringPattern(s string) bool {

	for i := 0; i < len(s)/2; i++ {
		flag := true
		j := i + 1

		for ; j+i < len(s); j += i + 1 {
			if s[:i+1] != s[j:j+i+1] {
				flag = false
			}
		}

		if flag && j >= len(s) {
			return true
		}
	}

	return false
}
