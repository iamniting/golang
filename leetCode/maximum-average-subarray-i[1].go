package main

func findMaxAverage(nums []int, k int) float64 {

	getSum := func(arr []int) (sum int) {
		for _, n := range arr {
			sum += n
		}
		return sum
	}

	sum := getSum(nums[:k])
	res := sum

	for i := 1; i <= len(nums)-k; i++ {
		sum += nums[i+k-1] - nums[i-1]
		if sum > res {
			res = sum
		}
	}

	return float64(res) / float64(k)
}
