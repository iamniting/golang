package main

func reverseVowels(s string) string {

	isVowel := func(c byte) bool {
		return (c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' ||
			c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u')
	}

	var res = []byte(s)
	i, j := 0, len(s)-1

	for i <= j {
		if !isVowel(res[i]) {
			i++
			continue
		}
		if !isVowel(res[j]) {
			j--
			continue
		}

		res[i], res[j] = res[j], res[i]
		i++
		j--
	}

	return string(res)
}
