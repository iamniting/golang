package main

func decrypt(code []int, k int) []int {

	res := make([]int, len(code))

	if k == 0 {
		return res
	}

	start, end, sum := 1, k, 0

	if k < 0 {
		start, end = len(code)-(-k), len(code)-1
	}

	for i := start; i <= end; i++ {
		sum += code[i]
	}

	for i := range code {
		res[i] = sum
		sum -= code[start%len(code)]
		start++
		end++
		sum += code[end%len(code)]
	}

	return res
}
