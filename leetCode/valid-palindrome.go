package main

func isPalindrome(s string) bool {
	str := ""

	for _, c := range s {
		c = c | ' '

		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			str += string(c)
		}
	}

	i, j := 0, len(str)-1

	for i <= j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}

	return true
}
