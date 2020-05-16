// https://leetcode.com/problems/array-partition-i
// Just sol to the problem, It does not include the I/O part

func arrayPairSum(nums []int) int {
	sort.Ints(nums)

	sum := 0
	for i := 0; i < len(nums); i = i + 2 {
		sum += nums[i]
	}

	return sum
}
