package main

func firstUniqChar(s string) int {
	freq := make([]int, 26)

	for _, c := range s {
		freq[c-'a']++
	}

	for i, c := range s {
		if freq[c-'a'] == 1 {
			return i
		}
	}

	return -1
}
