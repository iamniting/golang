package main

import "strings"

func reformatNumber(number string) string {

	var res string

	number = strings.ReplaceAll(number, "-", "")
	number = strings.ReplaceAll(number, " ", "")

	lastGroup := len(number) % 3

	for i := 0; i < len(number); i++ {

		if lastGroup == 2 && i+2 == len(number) {
			res += number[i:]
			break
		} else if lastGroup == 1 && i+4 == len(number) {
			res += number[i:i+2] + "-" + number[i+2:]
			break
		} else {
			res += number[i : i+3]
			i += 2
			if i+1 != len(number) {
				res += "-"
			}
		}
	}

	return res
}
