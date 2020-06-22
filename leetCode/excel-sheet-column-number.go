package main

func titleToNumber(s string) int {
	res := 0
	for _, c := range s {
		res = res*26 + int(c-'A'+1)
	}
	return res
}
