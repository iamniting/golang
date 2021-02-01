package main

import "math"

func minOperations(s string) int {

	var changeReqBaseZero, changeReqBaseOne int
	baseZero, baseOne := '0', '1'

	for i, c := range s {
		if i&1 == 0 && c != baseZero {
			changeReqBaseZero++
		} else if i&1 == 1 && c != baseOne {
			changeReqBaseZero++
		}

		if i&1 == 0 && c != baseOne {
			changeReqBaseOne++
		} else if i&1 == 1 && c != baseZero {
			changeReqBaseOne++
		}
	}

	return int(math.Min(float64(changeReqBaseZero), float64(changeReqBaseOne)))
}
