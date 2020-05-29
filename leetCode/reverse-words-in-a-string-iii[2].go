package main

func reverseWords(s string) string {
	// helper func to reverse a string
	reverse := func(word string, res []byte, index int) {
		for i := len(word) - 1; i >= 0; i-- {
			res[index] = word[i]
			index++
		}
	}

	res, index := []byte(s), 0
	for i, c := range s {
		if c == ' ' {
			reverse(s[index:i], res, index)
			index = i + 1
			// if reaches at end
		} else if i == len(s)-1 {
			reverse(s[index:], res, index)
		}
	}
	return string(res)
}
