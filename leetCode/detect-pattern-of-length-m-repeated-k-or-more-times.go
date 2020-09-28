package main

func containsPattern(arr []int, m int, k int) bool {

	for i := 0; i <= len(arr)-m; i++ {
		var count int
		for j := i; j <= len(arr)-m; j += m {
			if !isSame(arr[i:i+m], arr[j:j+m]) {
				break
			}

			count++
			if count == k {
				return true
			}
		}
	}

	return false
}

func isSame(a, b []int) bool {

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
