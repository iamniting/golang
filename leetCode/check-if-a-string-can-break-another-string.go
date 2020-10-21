package main

import "sort"

func checkIfCanBreak(s1 string, s2 string) bool {

	sortString := func(s string) string {
		r := []rune(s)
		sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
		return string(r)
	}

	if len(s1) != len(s2) {
		return false
	}

	s1 = sortString(s1)
	s2 = sortString(s2)

	var isLess, isGreat bool

	for i := range s1 {
		if s1[i] < s2[i] {
			isLess = true
		} else if s1[i] > s2[i] {
			isGreat = true
		}
	}

	return isLess != isGreat || (!isLess && !isGreat)
}
