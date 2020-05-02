// https://leetcode.com/problems/count-primes
// Just sol to the problem, It does not include the I/O part

func countPrimes(n int) int {
	var count int
	notPrime := make([]bool, n)

	for i := 2; i < n; i++ {
		if notPrime[i] {
			continue
		}
		count++

        /*
        mark all multiples of i,
        they are not prime, because they are divisible by i
        */
		for j := i * 2; j < n; j += i {
			notPrime[j] = true
		}
	}
	return count
}
