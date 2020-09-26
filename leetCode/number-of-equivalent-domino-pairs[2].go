package main

func numEquivDominoPairs(dominoes [][]int) int {

	var res int
	m := map[[2]int]int{}

	for i := range dominoes {

		a, b := dominoes[i][0], dominoes[i][1]

		if a > b {
			a, b = b, a
		}

		if count, ok := m[[2]int{a, b}]; ok {
			res += count
		}

		m[[2]int{a, b}]++
	}
	return res
}
