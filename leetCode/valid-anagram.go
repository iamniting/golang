package main

func isAnagram(s string, t string) bool {

	m1 := map[rune]int{}
	m2 := map[rune]int{}

	for _, c := range s {
		m1[c]++
	}

	for _, c := range t {
		m2[c]++
	}

	if len(m1) != len(m2) {
		return false
	}

	for k := range m1 {
		if _, ok := m2[k]; !ok || m1[k] != m2[k] {
			return false
		}
	}

	return true
}
