package main

func findJudge(N int, trust [][]int) int {

	var trustme [1001]int

	for _, t := range trust {
		trustme[t[0]]--
		trustme[t[1]]++
	}

	for i := 1; i <= N; i++ {
		if trustme[i] == N-1 {
			return i
		}
	}

	return -1
}
