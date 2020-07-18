package main

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	// if any of below true then not overlapping
	return !(rec1[0] >= rec2[2] || // rec2 is left side of rec1
		rec2[0] >= rec1[2] || // rec2 is right side of rec1
		rec1[1] >= rec2[3] || // rec2 is down of rec1
		rec2[1] >= rec1[3]) // rec2 is up of rec1
}
