package main

func diStringMatch(S string) []int {
	D, I := len(S), 0
	res := make([]int, len(S)+1)

	for i, c := range S + S[len(S)-1:] {
		if c == 'I' {
			res[i] = I
			I++
		} else if c == 'D' {
			res[i] = D
			D--
		}
	}
	return res
}
