package main

import "strconv"

func convertToBase7(num int) string {

	var res string
	var neg string

	if num == 0 {
		return "0"
	}

	if num < 0 {
		num = -num
		neg = "-"
	}

	for num > 0 {
		res = strconv.Itoa(num%7) + res
		num /= 7
	}

	return neg + res
}
