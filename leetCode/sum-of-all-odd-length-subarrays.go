package main

func sumOddLengthSubarrays(arr []int) int {

	calculateSum := func(a []int) int {
		var sum int
		for _, n := range a {
			sum += n
		}
		return sum
	}

	var res int

	// make pairs of odd len
	for i := 1; i <= len(arr); i += 2 {
		for j := 0; j <= len(arr)-i; j++ {
			res += calculateSum(arr[j : j+i])
		}
	}

	return res
}
