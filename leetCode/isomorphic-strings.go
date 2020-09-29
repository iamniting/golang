package main

func isIsomorphic(s string, t string) bool {

	m1 := map[byte]byte{}
	m2 := map[byte]byte{}

	for i := range s {

		if c, ok := m1[s[i]]; !ok {
			m1[s[i]] = t[i]
		} else if c != t[i] {
			return false
		}

		if c, ok := m2[t[i]]; !ok {
			m2[t[i]] = s[i]
		} else if c != s[i] {
			return false
		}
	}

	return true
}
