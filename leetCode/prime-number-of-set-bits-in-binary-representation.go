package main

func countPrimeSetBits(L int, R int) int {

	countBits := func(n int) int {
		res := 0
		for ; n > 0; n >>= 1 {
			res += n & 1
		}
		return res
	}

	isPrime := func(n int) bool {
		if n == 1 {
			return false
		}
		for i := 2; i <= n/2; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}

	res := 0
	for i := L; i <= R; i++ {
		bitsCount := countBits(i)
		if isPrime(bitsCount) {
			res++
		}
	}

	return res
}
