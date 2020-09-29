package main

func isIsomorphic(s string, t string) bool {

	m1 := [128]byte{}
	m2 := [128]byte{}

	for i := range s {
		if m1[s[i]] == 0 {
			m1[s[i]] = t[i]
		}

		if m2[t[i]] == 0 {
			m2[t[i]] = s[i]
		}

		if m1[s[i]] != t[i] {
			return false
		}

		if m2[t[i]] != s[i] {
			return false
		}
	}

	return true
}
