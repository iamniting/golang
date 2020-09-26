package main

func numEquivDominoPairs(dominoes [][]int) int {

	var res int

	for i := range dominoes {
		for j := i + 1; j < len(dominoes); j++ {
			a, b, c, d := dominoes[i][0], dominoes[i][1], dominoes[j][0], dominoes[j][1]
			if (a == c && b == d) || (a == d && b == c) {
				res++
			}
		}
	}
	return res
}
