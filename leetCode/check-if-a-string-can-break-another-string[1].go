package main

func checkIfCanBreak(s1 string, s2 string) bool {

	var a1, a2 [26]int

	for i := range s1 {
		a1[s1[i]-'a']++
		a2[s2[i]-'a']++
	}

	return helper(a1, a2) || helper(a2, a1)
}

func helper(a1, a2 [26]int) bool {
	var count int

	for i := range a1 {
		count += a1[i] - a2[i]
		if count < 0 {
			return false
		}
	}

	return true
}
