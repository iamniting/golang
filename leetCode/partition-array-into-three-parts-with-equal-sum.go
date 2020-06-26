package main

func canThreePartsEqualSum(A []int) bool {
	var totalsum, parts, requiredSum, sum int

	for _, n := range A {
		totalsum += n
	}

	if totalsum%3 != 0 {
		return false
	}

	requiredSum = totalsum / 3

	for _, n := range A {
		sum += n
		if sum == requiredSum {
			sum = 0
			totalsum -= requiredSum
			parts++
		}
		if parts == 3 {
			break
		}
	}

	return parts == 3 && totalsum == 0
}
