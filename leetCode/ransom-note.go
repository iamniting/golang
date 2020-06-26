package main

func canConstruct(ransomNote string, magazine string) bool {

	freq := make([]int, 26)

	for _, c := range ransomNote {
		freq[c-'a']++
	}

	for _, c := range magazine {
		freq[c-'a']--
	}

	for _, v := range freq {
		if v > 0 {
			return false
		}
	}

	return true
}
