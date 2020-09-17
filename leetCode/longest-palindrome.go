package main

func longestPalindrome(s string) int {

	var res int
	freq := map[byte]int{}

	for i := range s {
		freq[s[i]]++
	}

	for _, count := range freq {
		// simillar to res += (count / 2) * 2
		res += (count >> 1) << 1
	}

	// if len of s is greater then res then add one in res
	// because one char can be added in string in middle to keep it palidrome
	if len(s) > res {
		return res + 1
	}

	return res
}
