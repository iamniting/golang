package main

func trailingZeroes(n int) int {

	/*
		trailing zeros can be made with 5 * 2 = 10
		so we need to find how many 5 and 2 are present in factorial of n
		and we can take min(5, 2)
		As we now that 2 are always more than 5 so there is no need to calculate
		how many 2 are present just calculate how many 5 are present
	*/

	var res int

	// 5 can be present only in multiple of 5
	for i := 5; i <= n; i += 5 {

		// calculate how many factors of 5 are present
		for j := i; j%5 == 0; j /= 5 {
			res++
		}
	}

	return res
}
