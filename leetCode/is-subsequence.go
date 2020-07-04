package main

func isSubsequence(s string, t string) bool {
	var curr int

	for i := 0; i < len(t) && curr < len(s); i++ {

		if t[i] == s[curr] {
			curr++
		}
	}

	return curr == len(s)
}
