package main

import (
	"strconv"
	"strings"
)

func multiply(num1 string, num2 string) string {

	if num1 == "0" || num2 == "0" {
		return "0"
	}

	var carry int
	var res string
	cal := make([]string, len(num2))

	// calculate multiplications and store it in cal
	for i := len(num2) - 1; i >= 0; i-- {
		index := len(num2) - i - 1
		cal[index] = strings.Repeat("0", index)

		carry = 0
		for j := len(num1) - 1; j >= 0; j-- {

			tmp := int(num2[i]-'0')*int(num1[j]-'0') + carry
			carry = tmp / 10
			cal[index] = string(tmp%10+'0') + cal[index]
		}

		if carry > 0 {
			cal[index] = strconv.Itoa(carry) + cal[index]
		}
	}

	// sum all multiplications
	for j := 0; j < len(cal[len(num2)-1]); j++ {

		sum := carry
		for i := 0; i < len(num2); i++ {

			if len(cal[i])-j-1 > -1 {
				sum += int(cal[i][len(cal[i])-j-1] - '0')
			}
		}

		res = string(sum%10+'0') + res
		carry = sum / 10
	}

	if carry > 0 {
		res = strconv.Itoa(carry) + res
	}

	return res
}
