package main

import "math"

func minOperations(s string) int {

	var res, changeReq int
	baseZero := '0'
	baseOne := '1'

	for i, c := range s {
		if i&1 == 0 && c != baseZero {
			changeReq++
		} else if i&1 == 1 && c != baseOne {
			changeReq++
		}
	}

	res, changeReq = changeReq, 0

	for i, c := range s {
		if i&1 == 0 && c != baseOne {
			changeReq++
		} else if i&1 == 1 && c != baseZero {
			changeReq++
		}
	}

	return int(math.Min(float64(res), float64(changeReq)))
}
