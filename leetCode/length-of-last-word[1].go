package main

func lengthOfLastWord(s string) int {

	res := 0
	flag := false

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && !flag {
			continue
		} else if s[i] != ' ' {
			flag = true
			res++
		} else {
			break
		}
	}

	return res
}
