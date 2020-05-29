package main

func sortArrayByParity(A []int) []int {

	even, odd := []int{}, []int{}

	for _, i := range A {
		if i&1 == 1 {
			odd = append(odd, i)
		} else {
			even = append(even, i)
		}
	}

	return append(even, odd...)
}
