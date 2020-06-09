package main

func rotatedDigits(N int) int {
	isValid := func(num int) bool {
		ret := false
		for ; num > 0; num /= 10 {
			switch num % 10 {
			case 2, 5, 6, 9:
				ret = true
			case 3, 4, 7:
				return false
			}
		}
		return ret
	}

	res := 0
	for i := 1; i <= N; i++ {
		if isValid(i) {
			res++
		}
	}
	return res
}
