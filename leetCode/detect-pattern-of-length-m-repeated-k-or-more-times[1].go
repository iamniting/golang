package main

func containsPattern(arr []int, m int, k int) bool {

	if m*k > len(arr) {
		return false
	}

	var count int

	for i := 0; i < len(arr)-m; i++ {
		if arr[i] != arr[i+m] {
			count = 0
			continue
		}

		count++

		if count == m*(k-1) {
			return true
		}
	}
	return false
}
