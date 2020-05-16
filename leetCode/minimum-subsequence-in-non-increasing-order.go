// https://leetcode.com/problems/minimum-subsequence-in-non-increasing-order
// Just sol to the problem, It does not include the I/O part

func minSubsequence(nums []int) []int {
	sort.Slice(nums, func(a, b int) bool { return nums[a] > nums[b] })

	sum := 0
	for _, n := range nums {
		sum += n
	}

	res_sum := 0
	res := []int{}

	for _, n := range nums {
		res_sum += n
		res = append(res, n)
		if res_sum > sum-res_sum {
			break
		}
	}
	return res
}
