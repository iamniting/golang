package main

func prefixesDivBy5(A []int) []bool {

	var res = make([]bool, len(A))
	var n int

	for i, bit := range A {
		n = ((n << 1) | bit) % 5
		res[i] = (n == 0)
	}
	return res
}
