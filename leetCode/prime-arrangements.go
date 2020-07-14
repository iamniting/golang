package main

func numPrimeArrangements(n int) int {

	var factorial func(n int) int
	factorial = func(n int) int {
		if n == 1 || n == 0 {
			return 1
		}
		return n * factorial(n-1) % (1e9 + 7)
	}

	var pcount int
	var primes = []int{
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41,
		43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
	}

	for _, p := range primes {
		if p <= n {
			pcount++
		}
	}

	return factorial(pcount) * factorial(n-pcount) % (1e9 + 7)
}
