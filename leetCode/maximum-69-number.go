package main

import "strconv"

func maximum69Number(num int) int {

	numStr := strconv.Itoa(num)

	for i, c := range numStr {
		if c == '6' {
			num, _ = strconv.Atoi(numStr[:i] + "9" + numStr[i+1:])
			return num
		}
	}
	return num
}
