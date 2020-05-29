package main

func numberOfSteps(num int) int {

	res := 0

	for num > 0 {
		// (num & 1) = 1 for the odd num which means if odd num +2 in res
		res += (num & 1) + 1
		num >>= 1
	}

	return res - 1
}
