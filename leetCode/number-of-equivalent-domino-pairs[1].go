package main

func numEquivDominoPairs(dominoes [][]int) int {

	var res int
	type domino struct {
		a, b int
	}
	m := map[domino]int{}

	for i := range dominoes {

		a, b := dominoes[i][0], dominoes[i][1]

		if a == b {
			m[domino{a, b}]++
			res += m[domino{a, b}] - 1
			continue
		}

		if count, ok := m[domino{a, b}]; ok {
			res += count
		}
		if count, ok := m[domino{b, a}]; ok {
			res += count
		}
		m[domino{a, b}]++

	}
	return res
}
