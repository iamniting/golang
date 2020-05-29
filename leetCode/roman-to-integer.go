package main

func romanToInt(s string) int {

	m := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	res := 0

	for i, c := range s {

		// if next char is greater than the current one
		if i < len(s)-1 && m[c] < m[rune(s[i+1])] {
			res -= m[c]
			continue
		}

		res += m[c]
	}
	return res
}
