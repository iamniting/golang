package main

func mySqrt(x int) int {

	lft, rgt := 1, x

	for lft <= rgt {
		mid := (lft + rgt) / 2

		if val := mid * mid; val == x {
			return mid
		} else if val > x && (mid-1)*(mid-1) < x {
			return mid - 1
		} else if val > x {
			rgt = mid - 1
		} else if val < x {
			lft = mid + 1
		}
	}

	return 0
}
