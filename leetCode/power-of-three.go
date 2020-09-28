package main

func isPowerOfThree(n int) bool {

	for power := 1; power <= n; power *= 3 {
		if power == n {
			return true
		}

	}

	return false
}
