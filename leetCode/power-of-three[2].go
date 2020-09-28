package main

func isPowerOfThree(n int) bool {

	/*
		Knowing the limitation of n, we can now deduce that the maximum value of n
		that is also a power of three is 1162261467. We calculate this as 3**19 = 1162261467

		Therefore, the possible values of n where we should return true are 3**0, 3**1, 3**2, ... 3**19.
		Since 3 is a prime number, the only divisors of 3**19 are 3**0, 3**1, 3**2, ... 3**19.
		therefore all we need to do is divide 3**19 by n. A remainder of 0 means n is a divisor of 3**19
		and therefore a power of three.

		(This method can apply on all prime nums)

		eg.
		3 % 3 = 0
		9 % 3 = 0
		9 % 9 = 0
		27 % 3 = 0
		27 % 9 = 0
		27 % 27 = 0
		81 % 3 = 0
		81 % 9 = 0
		81 % 27 = 0
		81 % 81 = 0
	*/
	return n > 0 && 1162261467%n == 0
}
