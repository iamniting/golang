package main

func countOdds(low int, high int) int {
	// if nums is even, the number of even and odd numbers in this range will be the same
	// if nums is odd, the number of even and odd numbers in this range will not be the same

	nums := high - low + 1

	// if nums is odd and low is odd
	if nums&1 == 1 && low&1 == 1 {
		return (nums / 2) + 1
	}

	// if nums is odd and low is even
	// if nums is even
	return nums / 2
}
