package main

func mostVisited(n int, rounds []int) []int {

	res, start, end := []int{}, rounds[0], rounds[len(rounds)-1]

	if start <= end {

		for i := start; i <= end; i++ {
			res = append(res, i)
		}

	} else {

		for i := 1; i <= end; i++ {
			res = append(res, i)
		}

		for i := start; i <= n; i++ {
			res = append(res, i)
		}
	}

	return res
}
