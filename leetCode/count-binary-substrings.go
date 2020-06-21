package main

func countBinarySubstrings(s string) int {
	prevCount, currCount := 0, 1
	res := 0

	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			currCount++
		} else {
			prevCount = currCount
			currCount = 1
		}

		if prevCount >= currCount {
			res++
		}
	}

	return res
}
