package main

func maxProduct(nums []int) int {

	n1, n2 := 0, 0

	for _, n := range nums {

		if n >= n1 {
			n1, n2 = n, n1
		}

		if n > n2 && n < n1 {
			n2 = n
		}
	}

	return (n1 - 1) * (n2 - 1)
}
