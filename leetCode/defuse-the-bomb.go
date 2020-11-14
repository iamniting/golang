package main

func decrypt(code []int, k int) []int {

	res := make([]int, len(code))

	if k == 0 {
		return res
	}

	for i := range code {
		if k > 0 {
			for j := i + 1; j <= i+k; j++ {
				// As the array is circular, use modulo to find the correct index.
				res[i] += code[j%len(code)]
			}
		} else {
			for j := i + k; j < i; j++ {
				// As the array is circular, use modulo to find the correct index.
				res[i] += code[(j+len(code))%len(code)]
			}
		}
	}

	return res
}
