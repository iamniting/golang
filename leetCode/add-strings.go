package main

import "strconv"

func addStrings(num1 string, num2 string) string {

	var res string
	var carry int
	x, y := len(num1)-1, len(num2)-1

	for x >= 0 || y >= 0 || carry != 0 {
		sum := carry

		if x >= 0 {
			sum += int(num1[x] - '0')
			x--
		}
		if y >= 0 {
			sum += int(num2[y] - '0')
			y--
		}

		carry = sum / 10
		res = strconv.Itoa(sum%10) + res
	}

	return res
}
