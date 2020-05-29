package main

func numJewelsInStones(J string, S string) int {

	res := 0
	type void struct{}
	m := make(map[rune]void)

	for _, j := range J {

		m[j] = void{}
	}

	for _, s := range S {

		if _, ok := m[s]; ok {
			res++
		}
	}

	return res
}
