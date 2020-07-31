package main

func largeGroupPositions(S string) [][]int {

	var res [][]int
	var i, j int

	for i < len(S)-1 {

		for j = i + 1; j < len(S) && S[i] == S[j]; j++ {
		}

		if j-i > 2 {
			res = append(res, []int{i, j - 1})
		}

		i = j
	}

	return res
}
