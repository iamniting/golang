package main

func getNoZeroIntegers(n int) []int {

	isNonZero := func(num int) bool {
		for ; num > 0; num /= 10 {
			if num%10 == 0 {
				return false
			}
		}
		return true
	}

	for i := 1; i < n; i++ {
		if isNonZero(i) && isNonZero(n-i) {
			return []int{i, n - i}
		}
	}

	return []int{}
}
