package main

func findErrorNums(nums []int) []int {

	arr := make([]bool, len(nums))
	res := []int{}

	for _, n := range nums {
		if arr[n-1] {
			res = append(res, n)
		}
		arr[n-1] = true
	}

	for i, flag := range arr {
		if !flag {
			res = append(res, i+1)
		}
	}

	return res
}
