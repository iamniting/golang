package main

func isPalindrome(x int) bool {
	tmp := x
	res := 0

	for tmp > 0 {
		mod := tmp % 10
		tmp = tmp / 10
		res = (res * 10) + mod
	}

	return x == res
}
