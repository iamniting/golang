package main

import "math"

func findLUSlength(a string, b string) int {

	if a == b {
		// if a & b is same there can not be any uncommon subsequence
		return -1
	} else if len(a) == len(b) {
		// if a & b are not same but len is same, we can consider len of
		// any of string, because the whole string can not be a
		// subsequence of another one as there len is same
		return len(a)
	} else if len(a) != len(b) {
		// if a & b len is not same, we can consider the big len as big
		// len can not be found in the small one
		return int(math.Max(float64(len(a)), float64(len(b))))
	}

	return 0
}
