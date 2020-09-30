package main

func checkIfExist(arr []int) bool {

	m := map[int]bool{}

	for _, n := range arr {
		if m[n*2] || m[n/2] && n&1 == 0 {
			return true
		}
		m[n] = true
	}

	return false
}
