package main

func findNumbers(nums []int) int {

	res := 0

	for _, n := range nums {

		count := 0

		for ; n > 0; n /= 10 {
			count++
		}

		if count&1 == 0 {
			res++
		}
	}

	return res
}
