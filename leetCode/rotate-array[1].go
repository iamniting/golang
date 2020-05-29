package main

func rotate(nums []int, k int) {
	l := len(nums)
	k = k % l

	for i := 0; i < k; i++ {
		tmp := nums[l-1]

		for j := l - 2; j >= 0; j-- {
			nums[j+1] = nums[j]
		}
		nums[0] = tmp
	}
}
