package main

func createTargetArray(nums []int, index []int) []int {

	res := []int{}

	for i := 0; i < len(nums); i++ {
		// if index is greater than len, means need to append element
		if len(res)-1 < index[i] {
			res = append(res, nums[i])
			// else if index lies b/w
		} else {
			res = append(
				res[:index[i]], append([]int{nums[i]}, res[index[i]:]...)...)
		}
	}
	return res
}
