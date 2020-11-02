package main

func canFormArray(arr []int, pieces [][]int) bool {

	m := make(map[int][]int, len(pieces))

	for _, p := range pieces {
		m[p[0]] = p[1:]
	}

	for i := 0; i < len(arr); {
		piece, ok := m[arr[i]]

		if !ok {
			return false
		}

		i++

		for _, p := range piece {
			if arr[i] != p {
				return false
			}
			i++
		}
	}

	return true
}
