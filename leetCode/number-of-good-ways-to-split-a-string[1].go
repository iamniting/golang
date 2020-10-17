package main

func numSplits(s string) int {

	var res, lcount, rcount int
	var l, r [26]int

	for i := range s {
		r[s[i]-'a']++
		if r[s[i]-'a'] == 1 {
			rcount++
		}
	}

	for i := range s {
		l[s[i]-'a']++
		r[s[i]-'a']--

		if l[s[i]-'a'] == 1 {
			lcount++
		}

		if r[s[i]-'a'] == 0 {
			rcount--
		}

		if lcount == rcount {
			res++
		}
	}

	return res
}
