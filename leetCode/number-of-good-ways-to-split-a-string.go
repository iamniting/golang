package main

func numSplits(s string) int {

	var res int
	l, r := map[byte]int{}, map[byte]int{}

	for i := range s {
		r[s[i]]++
	}

	for i := range s {
		l[s[i]]++
		r[s[i]]--

		if r[s[i]] == 0 {
			delete(r, s[i])
		}

		if len(l) == len(r) {
			res++
		}
	}

	return res
}
