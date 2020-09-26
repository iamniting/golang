package main

import "fmt"

func readBinaryWatch(num int) []string {

	var res []string

	for h := 0; h < 12; h++ {
		for m := 0; m < 60; m++ {
			if countBinary(h)+countBinary(m) == num {
				res = append(res, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}

	return res
}

func countBinary(num int) int {
	var res int

	for ; num > 0; num >>= 1 {
		res += num & 1
	}

	return res
}
