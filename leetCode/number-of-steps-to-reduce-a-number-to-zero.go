package main

func numberOfSteps(num int) int {

	res := 0

	for num > 0 {
		if num&1 == 1 {
			num--
		} else {
			num >>= 1
		}
		res++
	}

	return res
}
